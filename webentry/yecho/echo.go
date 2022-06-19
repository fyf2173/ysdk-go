package yecho

import (
	"context"
	"github.com/fyf2173/ysdk-go/webentry"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

type EchoServer struct {
	srv *echo.Echo
}

func NewEchoServer() *EchoServer {
	return &EchoServer{
		srv: echo.New(),
	}
}

// Start 启动echo服务，并执行收尾工作
func (es *EchoServer) Start(serverAddr string, wipesFn ...func()) error {
	// 启动服务
	go func() {
		_ = es.srv.Start(serverAddr)
	}()

	serverContext, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
	for _, fn := range wipesFn {
		fn()
	}
	return es.srv.Shutdown(serverContext)
}

// RegisterHandler 注册路由方法
func (es *EchoServer) RegisterHandler(fn func(*echo.Echo) http.Handler) {
	es.srv.Server.Handler = fn(es.srv)
	return
}

// RegisterHandlerWithEntry 注册路由路口方法
func (es *EchoServer) RegisterHandlerWithEntry(entries ...*webentry.Entry[echo.HandlerFunc]) {
	for _, v := range entries {
		es.srv.Add(v.Method, v.Path, v.Handler)
	}
	return
}

// RegisterHandlerWithEntryGroup 分组注册路由入口方法
func (es *EchoServer) RegisterHandlerWithEntryGroup(g *echo.Group, entries ...*webentry.Entry[echo.HandlerFunc]) {
	for _, v := range entries {
		g.Add(v.Method, v.Path, v.Handler)
	}
	return
}

// GroupHandler 分组处理
func (es *EchoServer) GroupHandler(prefix string, mw ...echo.MiddlewareFunc) *echo.Group {
	return es.srv.Group(prefix, mw...)
}
