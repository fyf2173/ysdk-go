package echoplus

import (
	"net/http"

	"github.com/fyf2173/ysdk-go/apisdk"
	"github.com/labstack/echo/v4"
)

func Post(path string, handler echo.HandlerFunc, desc string) *apisdk.Route {
	return apisdk.NewRoute(http.MethodPost, path, apisdk.OptHandler(handler), apisdk.OptEcho, apisdk.OptDesc(desc))
}

func Get(path string, handler echo.HandlerFunc, desc string) *apisdk.Route {
	return apisdk.NewRoute(http.MethodGet, path, apisdk.OptHandler(handler), apisdk.OptEcho, apisdk.OptDesc(desc))
}

func Put(path string, handler echo.HandlerFunc, desc string) *apisdk.Route {
	return apisdk.NewRoute(http.MethodPut, path, apisdk.OptHandler(handler), apisdk.OptEcho, apisdk.OptDesc(desc))
}

func Delete(path string, handler echo.HandlerFunc, desc string) *apisdk.Route {
	return apisdk.NewRoute(http.MethodDelete, path, apisdk.OptHandler(handler), apisdk.OptEcho, apisdk.OptDesc(desc))
}

func Patch(path string, handler echo.HandlerFunc, desc string) *apisdk.Route {
	return apisdk.NewRoute(http.MethodPatch, path, apisdk.OptHandler(handler), apisdk.OptEcho, apisdk.OptDesc(desc))
}

func Options(path string, handler echo.HandlerFunc, desc string) *apisdk.Route {
	return apisdk.NewRoute(http.MethodOptions, path, apisdk.OptHandler(handler), apisdk.OptEcho, apisdk.OptDesc(desc))
}
