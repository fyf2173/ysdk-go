/**
 * @Author: entere@126.com
 * @Description:
 * @File:  handler
 * @Version: 1.0.0
 * @Date: 2020/5/20 22:26
 */

package handler

import (
	"image"
	"sync"
)

// Context Context
type Context struct {
	wg         sync.WaitGroup
	PngCarrier *image.RGBA
}

func NewContext(pngCarrier *image.RGBA) *Context {
	return &Context{wg: sync.WaitGroup{}, PngCarrier: pngCarrier}
}

// IHandler 处理
type IHandler interface {
	Do(c *Context)
	SetNext(h IHandler) IHandler
	Run(c *Context) error
}

// Next 抽象出来的 可被合成复用的结构体
type Next struct {
	nextHandler IHandler
}

// SetNext 实现好的 可被复用的SetNext方法
// 例如 nullHandler.SetNext(argumentsHandler).SetNext(signHandler).SetNext(frequentHandler)
func (n *Next) SetNext(h IHandler) IHandler {
	n.nextHandler = h
	return h
}

// Run 执行
func (n *Next) Run(c *Context) (err error) {
	if n.nextHandler == nil {
		return
	}
	(n.nextHandler).Do(c)
	return (n.nextHandler).Run(c)
}

func NewNullHandler() *NullHandler {
	return &NullHandler{}
}

// NullHandler 空Handler
type NullHandler struct {
	Next
}

func (h *NullHandler) Do(c *Context) (err error) {
	return
}
