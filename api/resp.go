package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type CommPageResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

type PageData struct {
	TotalCount int         `json:"total_count"`
	Items      interface{} `json:"items"`
}

func ApiReturn(args ...interface{}) CommPageResp {
	var (
		code = 0
		msg  = "OK"
		data interface{}
	)
	if len(args) >= 2 && args[0] != 0 {
		msg = args[1].(string)
	}
	if len(args) >= 3 {
		data = args[2]
	}
	return CommPageResp{code, msg, data}
}

func EchoReturn(code int, msg string, data interface{}) *CommPageResp {
	var ar = &CommPageResp{}
	ar.Code = code
	ar.Msg = msg
	ar.Data = data
	return ar
}

// EchoSuccess echo success response
func EchoSuccess(ctx echo.Context, data ...interface{}) error {
	var tmpData interface{}
	if len(data) > 0 {
		tmpData = data[0]
	}
	resp := EchoReturn(0, "success", tmpData)
	return ctx.JSON(http.StatusOK, resp)
}

func EchoError(ctx echo.Context, err error) error {
	resp := EchoReturn(1, err.Error(), nil)
	return ctx.JSON(http.StatusOK, resp)
}
