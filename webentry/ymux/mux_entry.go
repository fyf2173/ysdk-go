package ymux

import (
	"net/http"

	"github.com/fyf2173/ysdk-go/webentry"
)

func Post(path string, handler http.HandlerFunc) *webentry.Entry[http.HandlerFunc] {
	return webentry.NewEntry("POST", path, handler)
}

func Get(path string, handler http.HandlerFunc) *webentry.Entry[http.HandlerFunc] {
	return webentry.NewEntry("GET", path, handler)
}

func Put(path string, handler http.HandlerFunc) *webentry.Entry[http.HandlerFunc] {
	return webentry.NewEntry("PUT", path, handler)
}

func Delete(path string, handler http.HandlerFunc) *webentry.Entry[http.HandlerFunc] {
	return webentry.NewEntry("DELETE", path, handler)
}

func Patch(path string, handler http.HandlerFunc) *webentry.Entry[http.HandlerFunc] {
	return webentry.NewEntry("PATCH", path, handler)
}

func Options(path string, handler http.HandlerFunc) *webentry.Entry[http.HandlerFunc] {
	return webentry.NewEntry("OPTIONS", path, handler)
}
