package http

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
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
		SetRequestHeader("Content-Type", "application/json")(req)
		b, _ := json.Marshal(params)
		SetRequestBody(bytes.NewBuffer(b))(req)
	}
}

func Request(method, link string, params interface{}, resp interface{}, ops ...Option) error {
	log.Printf(">>>> 开始请求【link=%s】，参数【%+v】", link, func() string {
		b, _ := json.Marshal(params)
		return string(b)
	}())
	req, err := http.NewRequest(method, link, nil)
	if err != nil {
		return err
	}
	if params != nil {
		ops = append(ops, JsonBody(params))
	}
	for _, op := range ops {
		op(req)
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("请求【link=%s】出错，err=%s", link, err)
		return err
	}

	// if response.StatusCode != 200 {
	// 	err := fmt.Errorf("接口【link=%+v】请求错误[status_code=%d]", link, response.StatusCode)
	// 	log.Printf("%s", err)
	// 	return err
	// }
	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("读取数据出错，err=%+v", err)
		return err
	}
	if resp == nil {
		return nil
	}
	if err := json.Unmarshal(bodyBytes, &resp); err != nil {
		log.Printf("解析数据出错，err=%+v", err)
		return err
	}
	if response.ContentLength <= DefaultRespSize && response.ContentLength > 0 {
		log.Printf("获取响应数据，body=%+v", string(bodyBytes))
	}
	log.Printf(">>>> 结束请求【%s】，响应数据【ContentLength=%d】 <<<<", link, response.ContentLength)
	return nil
}
