package util

import (
	"context"
	"sync"
)

// GraceShutdown 优雅停止
type GraceShutdown struct {
	cancelFunc context.CancelFunc
	Context    context.Context
	c          chan struct{}
	mux        sync.Mutex
}

// NewGraceShutdown 初始化
func NewGraceShutdown() *GraceShutdown {
	h := &GraceShutdown{}
	h.Context, h.cancelFunc = context.WithCancel(context.Background())
	h.c = make(chan struct{})
	return h
}

// Done 停止成功
func (h *GraceShutdown) Done() {
	h.mux.Lock()
	defer h.mux.Unlock()
	if h.c != nil {
		close(h.c)
		h.c = nil
	}
}

// Shutdown 同步停止
func (h *GraceShutdown) Shutdown() {
	h.cancelFunc()
	h.mux.Lock()
	c := h.c
	h.mux.Unlock()
	if c != nil {
		<-c
	}
}
