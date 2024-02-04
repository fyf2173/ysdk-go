/**
 * @Author: entere@126.com
 * @Description:
 * @File:  background_handler.go
 * @Version: 1.0.0
 * @Date: 2020/5/21 12:31
 */

package handler

import (
	"fmt"
	"image"

	"github.com/fyf2173/ysdk-go/xposter/core"
)

// BackgroundHandler 背景图
type BackgroundRemoteHandler struct {
	// 合成复用Next
	Next
	X      int
	Y      int
	Width  int
	Height int
	URL    string
}

// Do 地址逻辑
func (h *BackgroundRemoteHandler) Do(c *Context) {
	//获取背景 必须是PNG图
	bgImage, err := core.GetResourceReader(h.URL, h.Width, h.Height)
	if err != nil {
		panic(fmt.Errorf("core.GetResourceReader err：%v", err))
	}
	bgPoint := image.Point{
		X: h.X,
		Y: h.Y,
	}
	core.MergeImage(c.PngCarrier, bgImage, bgImage.Bounds().Min.Sub(bgPoint))
	return
}
