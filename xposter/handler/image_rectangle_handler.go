/**
 * @Author: entere@126.com
 * @Description:
 * @File:  image_local_handler
 * @Version: 1.0.0
 * @Date: 2020/5/22 08:51
 */

package handler

import (
	"github.com/fyf2173/ysdk-go/xposter/core"
	"image"
)

// ImageRectangleHandler 设置矩形图片
type ImageRectangleHandler struct {
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
func (h *ImageRectangleHandler) Do(c *Context) {
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		request := core.Rectangle{
			R:      h.R,
			G:      h.G,
			B:      h.B,
			A:      h.A,
			Width:  h.Width,
			Height: h.Height,
		}

		im, _ := core.ImageRectangle(request)

		srcPoint := image.Point{
			X: h.X,
			Y: h.Y,
		}
		core.MergeImage(c.PngCarrier, im, im.Bounds().Min.Sub(srcPoint))
		return
	}()
}
