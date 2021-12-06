package api

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
