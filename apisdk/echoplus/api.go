package echoplus

import (
	"net/http"

	"github.com/fyf2173/ysdk-go/apisdk"
	"github.com/fyf2173/ysdk-go/xctx"
	"github.com/labstack/echo/v4"
)

func ExitError(ctx echo.Context, err error) error {
	var ar = &apisdk.CommResp{Code: 1, Msg: err.Error(), Data: nil}
	ar.RequestId = xctx.CtxId(xctx.Wrap(ctx.Request().Context()))
	return ctx.JSON(http.StatusOK, ar)
}

func ExitSuccess(ctx echo.Context, data interface{}) error {
	var ar = &apisdk.CommResp{Code: 0, Msg: "ok", Data: data}
	ar.RequestId = xctx.CtxId(xctx.Wrap(ctx.Request().Context()))
	return ctx.JSON(http.StatusOK, ar)
}

func ExitSuccessPage(ctx echo.Context, totalCount int, data interface{}) error {
	var ar = &apisdk.CommResp{
		Code: 0,
		Msg:  "ok",
		Data: &apisdk.PageData{
			TotalCount: int(totalCount),
			Items:      data,
		},
	}
	ar.RequestId = xctx.CtxId(xctx.Wrap(ctx.Request().Context()))
	return ctx.JSON(http.StatusOK, ar)
}
