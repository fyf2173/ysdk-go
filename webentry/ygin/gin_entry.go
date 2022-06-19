package ygin

import (
	"github.com/fyf2173/ysdk-go/webentry"
	"github.com/gin-gonic/gin"
)

func Post(path string, handler gin.HandlerFunc) *webentry.Entry[gin.HandlerFunc] {
	return webentry.NewEntry("POST", path, handler)
}

func Get(path string, handler gin.HandlerFunc) *webentry.Entry[gin.HandlerFunc] {
	return webentry.NewEntry("GET", path, handler)
}

func Put(path string, handler gin.HandlerFunc) *webentry.Entry[gin.HandlerFunc] {
	return webentry.NewEntry("PUT", path, handler)
}

func Delete(path string, handler gin.HandlerFunc) *webentry.Entry[gin.HandlerFunc] {
	return webentry.NewEntry("DELETE", path, handler)
}

func Patch(path string, handler gin.HandlerFunc) *webentry.Entry[gin.HandlerFunc] {
	return webentry.NewEntry("PATCH", path, handler)
}

func Options(path string, handler gin.HandlerFunc) *webentry.Entry[gin.HandlerFunc] {
	return webentry.NewEntry("OPTIONS", path, handler)
}
