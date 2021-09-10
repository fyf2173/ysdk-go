package api

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

func (es *EchoServer) Start(serverAddr string) {
	// 启动服务
	go func() {
		es.srv.Start(serverAddr)
	}()

	serverContext, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
	es.srv.Shutdown(serverContext)
}

func (es *EchoServer) RegisterHandler(fn func(*echo.Echo)) {
	fn(es.srv)
}
