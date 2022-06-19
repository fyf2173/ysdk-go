package ygin

import (
	"context"
	"github.com/fyf2173/ysdk-go/webentry"
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
func (gs *GinServer) RegisterHandler(fn func(*gin.Engine)) {
	fn(gs.srv)
	return
}

func (gs *GinServer) RegisterHandlerWithEntry(entries ...*webentry.Entry[gin.HandlerFunc]) {
	for _, v := range entries {
		gs.srv.RouterGroup.Handle(v.Method, v.Path, v.Handler)
	}
	return
}

// RegisterHandlerWithEntryGroup 分组注册路由入口方法
func (gs *GinServer) RegisterHandlerWithEntryGroup(g *gin.RouterGroup, entries ...*webentry.Entry[gin.HandlerFunc]) {
	for _, v := range entries {
		g.Handle(v.Method, v.Path, v.Handler)
	}
	return
}

// GroupHandler 分组处理
func (gs *GinServer) GroupHandler(prefix string, mw ...gin.HandlerFunc) *gin.RouterGroup {
	g := gs.srv.Group(prefix)
	g.Use(mw...)
	return g
}
