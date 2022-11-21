package ginplus

import (
	"net/http"

	"github.com/fyf2173/ysdk-go/apisdk"
	"github.com/gin-gonic/gin"
)

func Post(path string, handler gin.HandlerFunc, desc string) *apisdk.Route {
	return apisdk.NewRoute(http.MethodPost, path, apisdk.OptHandler(handler), apisdk.OptGin, apisdk.OptDesc(desc))
}

func Get(path string, handler gin.HandlerFunc, desc string) *apisdk.Route {
	return apisdk.NewRoute(http.MethodGet, path, apisdk.OptHandler(handler), apisdk.OptGin, apisdk.OptDesc(desc))
}

func Put(path string, handler gin.HandlerFunc, desc string) *apisdk.Route {
	return apisdk.NewRoute(http.MethodPut, path, apisdk.OptHandler(handler), apisdk.OptGin, apisdk.OptDesc(desc))
}

func Delete(path string, handler gin.HandlerFunc, desc string) *apisdk.Route {
	return apisdk.NewRoute(http.MethodDelete, path, apisdk.OptHandler(handler), apisdk.OptGin, apisdk.OptDesc(desc))
}

func Patch(path string, handler gin.HandlerFunc, desc string) *apisdk.Route {
	return apisdk.NewRoute(http.MethodPatch, path, apisdk.OptHandler(handler), apisdk.OptGin, apisdk.OptDesc(desc))
}

func Options(path string, handler gin.HandlerFunc, desc string) *apisdk.Route {
	return apisdk.NewRoute(http.MethodOptions, path, apisdk.OptHandler(handler), apisdk.OptGin, apisdk.OptDesc(desc))
}
