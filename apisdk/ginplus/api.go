package ginplus

import (
	"net/http"

	"github.com/fyf2173/ysdk-go/apisdk"
	"github.com/fyf2173/ysdk-go/xctx"

	"github.com/gin-gonic/gin"
)

func ExitError(ctx *gin.Context, err error) {
	var ar = &apisdk.CommResp{Code: 1, Msg: err.Error(), Data: nil}
	ar.RequestId = xctx.CtxId(xctx.Wrap(ctx))
	ctx.JSON(http.StatusOK, ar)
	return
}

func ExitSuccess(ctx *gin.Context, data interface{}) {
	var ar = &apisdk.CommResp{Code: 0, Msg: "ok", Data: data}
	ar.RequestId = xctx.CtxId(xctx.Wrap(ctx))
	ctx.JSON(http.StatusOK, ar)
	return
}

func ExitSuccessPage(ctx *gin.Context, totalCount int, data interface{}) {
	var ar = &apisdk.CommResp{
		Code: 0,
		Msg:  "ok",
		Data: &apisdk.PageData{
			TotalCount: int(totalCount),
			Items:      data,
		},
	}
	ar.RequestId = xctx.CtxId(xctx.Wrap(ctx))
	ctx.JSON(http.StatusOK, ar)
	return
}
