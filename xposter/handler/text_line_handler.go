/**
 * @Author: entere@126.com
 * @Description:
 * @File:  background_handler.go
 * @Version: 1.0.0
 * @Date: 2020/5/21 12:31
 */

package handler

import (
	"image/color"
)

// TextLineHandler 文本
type TextLineHandler struct {
	// 合成复用Next
	Next
	X      int
	Y      int
	R      uint8
	G      uint8
	B      uint8
	A      uint8
	Width  int
	Height int
}

// Do 地址逻辑
func (h *TextLineHandler) Do(c *Context) {
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()

		for i := h.X; i < h.X+h.Width; i++ {
			for j := h.Y; j < h.Y+h.Height; j++ {
				c.PngCarrier.Set(i, j, color.RGBA{R: h.R, G: h.G, B: h.B, A: h.A})
			}
		}
		return
	}()
}
