package webentry

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

// Entry 入口方法
type Entry[T echo.HandlerFunc | gin.HandlerFunc] struct {
	Group   string
	Method  string
	Path    string
	Handler T
}

func NewEntry[T echo.HandlerFunc | gin.HandlerFunc](method, path string, handler T) *Entry[T] {
	return &Entry[T]{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
}

func Post[T echo.HandlerFunc | gin.HandlerFunc](path string, handler T) *Entry[T] {
	return NewEntry("POST", path, handler)
}

func Get[T echo.HandlerFunc | gin.HandlerFunc](path string, handler T) *Entry[T] {
	return NewEntry("GET", path, handler)
}

func Put[T echo.HandlerFunc | gin.HandlerFunc](path string, handler T) *Entry[T] {
	return NewEntry("PUT", path, handler)
}

func Delete[T echo.HandlerFunc | gin.HandlerFunc](path string, handler T) *Entry[T] {
	return NewEntry("DELETE", path, handler)
}

func Patch[T echo.HandlerFunc | gin.HandlerFunc](path string, handler T) *Entry[T] {
	return NewEntry("PATCH", path, handler)
}

func Options[T echo.HandlerFunc | gin.HandlerFunc](path string, handler T) *Entry[T] {
	return NewEntry("OPTIONS", path, handler)
}
