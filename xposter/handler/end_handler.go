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
	"sync"

	"github.com/fyf2173/ysdk-go/xposter/core"
)

// EndHandler 结束，写在最后，把图片合并到一张图上
type EndHandler struct {
	// 合成复用Next
	Next
	Output string // "/tmp/xxx.png"
}

// Do 地址逻辑
func (h *EndHandler) Do(c *Context, wg *sync.WaitGroup) (err error) {
	wg.Wait()

	// 新建文件载体
	//fileName := "poster-" + core.RandString(20) + ".png"
	merged, err := core.NewMerged(h.Output)
	if err != nil {
		fmt.Errorf("core.NewMerged err：%v", err)
	}
	// 合并
	err = core.Merge(c.PngCarrier, merged)
	if err != nil {
		fmt.Errorf("core.Merge err：%v", err)
	}
	return
}
