/**
 * @Author: entere@126.com
 * @Description:
 * @File:  background_handler.go
 * @Version: 1.0.0
 * @Date: 2020/5/21 12:31
 */

package handler

import (
	"bytes"
	"image/png"
	"os"
)

// BufferEndHandler 把图片输出为二进制
type BufferEndHandler struct {
	Next
	Output *bytes.Buffer
}

func (h BufferEndHandler) Do(c *Context) error {
	c.wg.Wait()
	return png.Encode(h.Output, c.PngCarrier)
}

// FileEndHandler 把图片保存到指定路径
type FileEndHandler struct {
	Next
	Filepath string
}

func (h FileEndHandler) Do(c *Context) {
	c.wg.Wait()
	nfi, err := os.Create(h.Filepath)
	if err != nil {
		panic(err)
	}
	defer nfi.Close()
	if err := png.Encode(nfi, c.PngCarrier); err != nil {
		panic(err)
	}
	return
}
