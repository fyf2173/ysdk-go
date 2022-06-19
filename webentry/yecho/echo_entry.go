package yecho

import (
	"github.com/fyf2173/ysdk-go/webentry"
	"github.com/labstack/echo/v4"
)

func Post(path string, handler echo.HandlerFunc) *webentry.Entry[echo.HandlerFunc] {
	return webentry.NewEntry("POST", path, handler)
}

func Get(path string, handler echo.HandlerFunc) *webentry.Entry[echo.HandlerFunc] {
	return webentry.NewEntry("GET", path, handler)
}

func Put(path string, handler echo.HandlerFunc) *webentry.Entry[echo.HandlerFunc] {
	return webentry.NewEntry("PUT", path, handler)
}

func Delete(path string, handler echo.HandlerFunc) *webentry.Entry[echo.HandlerFunc] {
	return webentry.NewEntry("DELETE", path, handler)
}

func Patch(path string, handler echo.HandlerFunc) *webentry.Entry[echo.HandlerFunc] {
	return webentry.NewEntry("PATCH", path, handler)
}

func Options(path string, handler echo.HandlerFunc) *webentry.Entry[echo.HandlerFunc] {
	return webentry.NewEntry("OPTIONS", path, handler)
}
