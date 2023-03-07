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
	"github.com/fyf2173/ysdk-go/xposter/circlemask"
	"github.com/fyf2173/ysdk-go/xposter/core"
	"image"
)

// ImageCircleRemoteHandler 根据URL地址设置圆形图片
type ImageCircleRemoteHandler struct {
	// 合成复用Next
	Next
	X      int
	Y      int
	Width  int
	Height int
	R      uint8
	G      uint8
	B      uint8
	A      uint8
	URL    string //http://xxx.png
}

// Do 地址逻辑
func (h *ImageCircleRemoteHandler) Do(c *Context) {
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		srcImage, err := core.GetResourceReader(h.URL, h.Width, h.Height)
		if err != nil {
			panic(fmt.Errorf("core.GetResourceReader err：%v", err))
		}

		// 算出图片的宽度和高试
		width := srcImage.Bounds().Max.X - srcImage.Bounds().Min.X
		hight := srcImage.Bounds().Max.Y - srcImage.Bounds().Min.Y

		//把头像转成Png,否则会有白底
		srcPng := core.NewPNG(0, 0, width, hight)
		core.MergeImage(srcPng, srcImage, srcImage.Bounds().Min)

		// 圆的直径以长边为准
		diameter := width
		if width > hight {
			diameter = hight
		}
		// 遮罩
		srcMask := circlemask.NewCircleMask(srcPng, image.Point{0, 0}, diameter, h.R, h.G, h.B, h.A)

		srcPoint := image.Point{
			X: h.X,
			Y: h.Y,
		}
		core.MergeImage(c.PngCarrier, srcMask, srcImage.Bounds().Min.Sub(srcPoint))
		return
	}()
}
