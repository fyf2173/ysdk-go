/**
 * @Author: entere@126.com
 * @Description:
 * @File:  image_local_handler
 * @Version: 1.0.0
 * @Date: 2020/5/22 08:51
 */

package handler

import (
	"fmt"
	"github.com/fyf2173/ysdk-go/xposter/core"
	"image"
)

// ImageRemoteHandler 根据URL地址设置图片
type ImageRemoteHandler struct {
	// 合成复用Next
	Next
	X      int
	Y      int
	Width  int
	Height int
	URL    string //http://xxx.png
}

// Do 地址逻辑
func (h *ImageRemoteHandler) Do(c *Context) {
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		srcImage, err := core.GetResourceReader(h.URL, h.Width, h.Height)
		if err != nil {
			panic(fmt.Errorf("core.GetResourceReader err：%v", err))
		}

		srcPoint := image.Point{
			X: h.X,
			Y: h.Y,
		}
		core.MergeImage(c.PngCarrier, srcImage, srcImage.Bounds().Min.Sub(srcPoint))
		return
	}()
}
