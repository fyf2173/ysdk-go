package ginplus

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

type GinServer struct {
	srv *gin.Engine
}

func NewGinServer() *GinServer {
	return &GinServer{
		srv: gin.New(),
	}
}

// Start 启动echo服务，并执行收尾工作
func (gs *GinServer) Start(serverAddr string, wipesFn ...func()) error {
	// 启动服务

	go func() {
		_ = gs.srv.Run(serverAddr)
	}()

	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit

	for _, fn := range wipesFn {
		fn()
	}

	return nil
}

// RegisterHandler 注册路由方法
func (gs *GinServer) RegisterHandler(fns ...func(*gin.Engine)) {
	for _, fn := range fns {
		fn(gs.srv)
	}
	return
}
