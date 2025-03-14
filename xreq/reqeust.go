package xreq

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/fyf2173/ysdk-go/xctx"
)

const (
	HeaderTraceId = "Trace-Id"
)

const DefaultRespSize = 500 * 1024 // unit Kb

type ReqOption func(*http.Request)

func SetRequestHeader(key, value string) ReqOption {
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

// SetContentType set request content type.
func SetContentType(contentType string) ReqOption {
	return func(request *http.Request) {
		request.Header.Set("Content-Type", contentType)
	}
}

func SetTraceId(ctx context.Context) ReqOption {
	return func(req *http.Request) {
		req.Header.Set(HeaderTraceId, xctx.CtxId(ctx))
	}
}

// SetRequestBody set request body.
func SetRequestBody(body io.Reader) ReqOption {
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

func JsonBody(params interface{}) ReqOption {
	return func(req *http.Request) {
		SetContentTypeJson(req)
		b, _ := json.Marshal(params)
		SetRequestBody(bytes.NewBuffer(b))(req)
	}
}

// FromdataBody submit multipart/form-data.
func FromdataBody(params map[string]string, files ...*os.File) ReqOption {
	return func(req *http.Request) {
		rb := &bytes.Buffer{} // create a new buffer
		w := multipart.NewWriter(rb)
		for _, fi := range files {
			fw, err := w.CreateFormFile("files", fi.Name()) // custom file name
			if err != nil {
				panic(err)
			}
			// copy file content to buffer
			if _, err := io.Copy(fw, fi); err != nil {
				panic(err)
			}
			fi.Close()
		}
		for k, v := range params {
			w.WriteField(k, v)
		}
		w.Close() // close writer, very important, otherwise the request will not be sent
		SetContentType(w.FormDataContentType())(req)
		SetRequestBody(rb)(req)
	}
}

// FormBody submit application/x-www-form-urlencoded.
func FormBody(params map[string]string) ReqOption {
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
