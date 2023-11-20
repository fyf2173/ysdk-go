package muxplus

import (
	"encoding/json"
	"net/http"

	"github.com/fyf2173/ysdk-go/apisdk"
	"github.com/fyf2173/ysdk-go/xctx"
)

func ExitError(r http.Request, w http.ResponseWriter, err error) error {
	var ar = &apisdk.CommResp{Code: 1, Msg: err.Error(), Data: nil}
	ar.RequestId = xctx.CtxId(xctx.Wrap(r.Context()))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(ar)
}

func ExitSuccess(r http.Request, w http.ResponseWriter, data interface{}) error {
	var ar = &apisdk.CommResp{Code: 0, Msg: "ok", Data: data}
	ar.RequestId = xctx.CtxId(xctx.Wrap(r.Context()))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(ar)
}

func ExitSuccessPage(r http.Request, w http.ResponseWriter, totalCount int, data interface{}) error {
	var ar = &apisdk.CommResp{
		Code: 0,
		Msg:  "ok",
		Data: &apisdk.PageData{
			TotalCount: int(totalCount),
			Items:      data,
		},
	}
	ar.RequestId = xctx.CtxId(xctx.Wrap(r.Context()))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(ar)
}
