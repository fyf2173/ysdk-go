package handler

import (
	"bytes"
	"fmt"
	"github.com/fyf2173/ysdk-go/xposter/core"
	"image"
	"sync"
)

// ImageBufferHandler 根据二进制内容设置图片
type ImageBufferHandler struct {
	Next
	X      int
	Y      int
	Width  int
	Height int
	Buf    *bytes.Reader
}

func (h *ImageBufferHandler) Do(c *Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		srcImage, _, err := image.Decode(h.Buf)
		if err != nil {
			panic(fmt.Errorf("png.Decode err：%v", err))
		}

		newImg := core.ResizeImage(srcImage, h.Width, h.Height)

		srcPoint := image.Point{
			X: h.X,
			Y: h.Y,
		}
		core.MergeImage(c.PngCarrier, newImg, newImg.Bounds().Min.Sub(srcPoint))
		return
	}()
}
