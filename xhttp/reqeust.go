package xhttp

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/fyf2173/ysdk-go/xctx"
	"github.com/fyf2173/ysdk-go/xlog"
)

const (
	HeaderTraceId = "Trace-Id"
)

const DefaultRespSize = 500 * 1024 // 单位Kb

type Option func(*http.Request)

func SetRequestHeader(key, value string) Option {
	return func(request *http.Request) {
		request.Header.Set(key, value)
	}
}

var (
	SetContentTypeJson     = SetContentType("application/json")
	SetContentTypeXml      = SetContentType("application/xml")
	SetContentTypeFormData = SetContentType("multipart/form-data")
	SetContentTypeForm     = SetContentType("application/x-www-form-urlencoded")
	SetContentTypeText     = SetContentType("text/plain; charset=utf-8")
)

// SetContentType 设置请求头Content-Type类型
func SetContentType(contentType string) Option {
	return func(request *http.Request) {
		request.Header.Set("Content-Type", contentType)
	}
}

func SetTraceId(ctx context.Context) Option {
	return func(req *http.Request) {
		req.Header.Set(HeaderTraceId, xctx.CtxId(ctx))
	}
}

// SetRequestBody 设置请求参数
func SetRequestBody(body io.Reader) Option {
	return func(req *http.Request) {
		rc, ok := body.(io.ReadCloser)
		if !ok && body != nil {
			rc = io.NopCloser(body)
		}
		req.Body = rc
		if body != nil {
			switch v := body.(type) {
			case *bytes.Buffer:
				req.ContentLength = int64(v.Len())
				buf := v.Bytes()
				req.GetBody = func() (io.ReadCloser, error) {
					r := bytes.NewReader(buf)
					return io.NopCloser(r), nil
				}
			case *bytes.Reader:
				req.ContentLength = int64(v.Len())
				snapshot := *v
				req.GetBody = func() (io.ReadCloser, error) {
					r := snapshot
					return io.NopCloser(&r), nil
				}
			case *strings.Reader:
				req.ContentLength = int64(v.Len())
				snapshot := *v
				req.GetBody = func() (io.ReadCloser, error) {
					r := snapshot
					return io.NopCloser(&r), nil
				}
			default:
				// This is where we'd set it to -1 (at least
				// if body != NoBody) to mean unknown, but
				// that broke people during the Go 1.8 testing
				// period. People depend on it being 0 I
				// guess. Maybe retry later. See Issue 18117.
			}
			// For client requests, Request.ContentLength of 0
			// means either actually 0, or unknown. The only way
			// to explicitly say that the ContentLength is zero is
			// to set the Body to nil. But turns out too much code
			// depends on NewRequest returning a non-nil Body,
			// so we use a well-known ReadCloser variable instead
			// and have the http package also treat that sentinel
			// variable to mean explicitly zero.
			if req.GetBody != nil && req.ContentLength == 0 {
				req.Body = http.NoBody
				req.GetBody = func() (io.ReadCloser, error) { return http.NoBody, nil }
			}
		}
	}
}

func JsonBody(params interface{}) Option {
	return func(req *http.Request) {
		SetContentTypeJson(req)
		b, _ := json.Marshal(params)
		SetRequestBody(bytes.NewBuffer(b))(req)
	}
}

// FromdataBody 文件表单提交(multipart/form-data)
func FromdataBody(params map[string]string, files ...*os.File) Option {
	return func(req *http.Request) {
		var rb = &bytes.Buffer{} // 创建一个buffer
		w := multipart.NewWriter(rb)
		for _, fi := range files {
			fw, err := w.CreateFormFile("files", fi.Name()) // 自定义文件名，发送文件流
			if err != nil {
				panic(err)
			}
			// 把文件内容，复制到fw中
			if _, err := io.Copy(fw, fi); err != nil {
				panic(err)
			}
			fi.Close()
		}
		for k, v := range params {
			w.WriteField(k, v)
		}
		w.Close() // 很重要，一定要关闭写入，不然服务端会报EOF错误，而且度不到数据
		SetContentType(w.FormDataContentType())(req)
		SetRequestBody(rb)(req)
	}
}

// FormBody 普通表单提交(application/x-www-form-urlencoded)
func FormBody(params map[string]string) Option {
	return func(req *http.Request) {
		var form url.Values
		for k, v := range params {
			form.Set(k, v)
		}
		SetContentTypeForm(req)
		SetRequestBody(bytes.NewBufferString(form.Encode()))(req)
	}
}

type IResponse interface {
	Unmarshal(src []byte, dst interface{}) error
}

type JsonResponse struct{}

func (jr *JsonResponse) Unmarshal(src []byte, dst interface{}) error {
	return json.Unmarshal(src, &dst)
}

type XmlResponse struct{}

func (xr *XmlResponse) Unmarshal(src []byte, dst interface{}) error {
	return xml.Unmarshal(src, &dst)
}

func Request(ctx context.Context, method, link string, params interface{}, resp IResponse, ops ...Option) error {
	xlog.Info(ctx, fmt.Sprintf(">>> 开始请求【[%s]link=%s】", method, link), slog.Any("params", params))
	req, err := http.NewRequest(method, link, nil)
	if err != nil {
		return err
	}
	ops = append(ops, SetTraceId(ctx))
	if params != nil {
		ops = append(ops, JsonBody(params))
	}
	for _, op := range ops {
		op(req)
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		xlog.Error(ctx, err, slog.String("method", method), slog.String("link", link))
		return err
	}

	if response.StatusCode != 200 {
		xlog.Info(ctx, fmt.Sprintf("[%s]%s:%d", method, link, response.StatusCode))
		return fmt.Errorf("errorstatus:%d", response.StatusCode)
	}
	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		xlog.Error(ctx, err)
		return err
	}
	if response.ContentLength <= DefaultRespSize && response.ContentLength > 0 {
		xlog.Info(ctx, "trace response", slog.String("response", string(bodyBytes)))
	}
	xlog.Info(ctx, fmt.Sprintf(">>> 结束请求[%s]%s", method, link), slog.Int64("content_length", response.ContentLength))
	if resp == nil {
		return nil
	}
	return resp.Unmarshal(bodyBytes, resp)
}
