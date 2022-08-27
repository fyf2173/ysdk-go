package ymux

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type MuxServer struct {
	srv *http.Server
}

func NewMuxServer() *MuxServer {
	return &MuxServer{
		srv: &http.Server{},
	}
}

// Start 启动echo服务，并执行收尾工作
func (ms *MuxServer) Start(serverAddr string, wipesFn ...func()) error {
	// 启动服务
	go func() {
		log.Println("⇨ http server started on ", serverAddr)

		ms.srv.Addr = serverAddr
		log.Println(ms.srv.ListenAndServe())
	}()

	serverContext, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
	for _, fn := range wipesFn {
		fn()
	}
	return ms.srv.Shutdown(serverContext)
}

// RegisterHandler 注册路由方法
func (ms *MuxServer) RegisterHandler(fns ...func(r *mux.Router)) {
	r := mux.NewRouter()
	for _, fn := range fns {
		fn(r)
	}
	ms.srv.Handler = r
	return
}
