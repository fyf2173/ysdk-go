package apisdk

import (
	"net"
	"net/http"
)

type CommPageReq struct {
	Page   int `json:"page" query:"page" form:"page"`
	Limit  int `json:"limit" query:"limit" form:"limit"`
	Offset int
}

func (cpr *CommPageReq) SetOffset() {
	if cpr.Page == 0 {
		cpr.Page = 1
	}
	if cpr.Limit == 0 {
		cpr.Limit = 12
	}
	cpr.Offset = (cpr.Page - 1) * cpr.Limit
}

type CommResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

type PageData struct {
	TotalCount int         `json:"total_count"`
	Items      interface{} `json:"items"`
}

func Exit(code int, msg string, data interface{}) *CommResp {
	var ar = &CommResp{}
	ar.Code = code
	ar.Msg = msg
	ar.Data = data
	return ar
}

func ExitError(err error) *CommResp {
	var ar = &CommResp{Code: 1, Msg: err.Error(), Data: nil}
	return ar
}

func ExitSuccess(data interface{}) *CommResp {
	var ar = &CommResp{Code: 0, Msg: "ok", Data: data}
	return ar
}

func ExitSuccessPage(totalCount int, data interface{}) *CommResp {
	var ar = &CommResp{
		Code: 0,
		Msg:  "ok",
		Data: &PageData{
			TotalCount: int(totalCount),
			Items:      data,
		},
	}
	return ar
}

func GetRealIp() string {
	var req = &http.Request{}
	var ip = req.Header.Get("X-Real-Ip")
	if ip == "" {
		ip = req.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip = req.RemoteAddr
	}
	if net.ParseIP(ip) != nil {
		return ip
	}
	host, _, _ := net.SplitHostPort(ip)
	return host
}
