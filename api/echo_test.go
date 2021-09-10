package api

import (
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestNewEcho(t *testing.T) {
	srv := NewEchoServer()
	srv.RegisterHandler(func(e *echo.Echo) {
		r := e.Group("/v1")
		r.GET("/alive", func(c echo.Context) error {
			return c.JSON(http.StatusOK, "Hello World")
		})
	})
	srv.RegisterHandler(func(e *echo.Echo) {
		r := e.Group("/v2")
		r.GET("/alive-h", func(c echo.Context) error {
			return c.JSON(http.StatusOK, "Hello World ~~~")
		})
	})
	srv.Start(":8080")
}
