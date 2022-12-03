package echoplus

import (
	"context"
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
func (es *EchoServer) RegisterHandler(fns ...func(*echo.Echo)) {
	for _, fn := range fns {
		fn(es.srv)
	}
	es.srv.Server.Handler = es.srv
	return
}
