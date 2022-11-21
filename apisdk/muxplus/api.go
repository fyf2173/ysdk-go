package muxplus

import (
	"encoding/json"
	"github.com/fyf2173/ysdk-go/apisdk"
	"net/http"
)

func ExitError(w http.ResponseWriter, err error) error {
	var ar = &apisdk.CommResp{Code: 1, Msg: err.Error(), Data: nil}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(ar)
}

func ExitSuccess(w http.ResponseWriter, data interface{}) error {
	var ar = &apisdk.CommResp{Code: 0, Msg: "ok", Data: data}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(ar)
}

func ExitSuccessPage(w http.ResponseWriter, totalCount int, data interface{}) error {
	var ar = &apisdk.CommResp{
		Code: 0,
		Msg:  "ok",
		Data: &apisdk.PageData{
			TotalCount: int(totalCount),
			Items:      data,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(ar)
}
