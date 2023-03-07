/**
 * @Author: entere@126.com
 * @Description:
 * @File:  handler_test
 * @Version: 1.0.0
 * @Date: 2020/5/20 22:39
 */

package handler

import (
	"fmt"
	"testing"

	"github.com/fyf2173/ysdk-go/xposter/core"
)

// TestNext_SetNext test
func TestNext_SetNext(t *testing.T) {
	ctx := NewContext(core.NewPNG(0, 0, 750, 1334))
	//绘制背景图
	backgroundHandler := &BackgroundHandler{
		X:    0,
		Y:    0,
		Path: "../assets/background.png",
	}
	//绘制圆形图像
	imageCircleHandler := &ImageCircleLocalHandler{
		X:    30,
		Y:    50,
		Path: "../assets/reward.png",
	}
	//绘制本地图像
	imageLocalHandler := &ImageLocalHandler{
		X:    30,
		Y:    400,
		Path: "../assets/reward.png",
	}

	//绘制二维码
	qrCodeHandler := &QRCodeHandler{
		X:   30,
		Y:   860,
		URL: "https://github.com/hitailang/poster",
	}
	//绘制文字
	textHandler1 := &TextHandler{
		Next:     Next{},
		X:        180,
		Y:        105,
		Size:     20,
		R:        255,
		G:        241,
		B:        250,
		Text:     "如果觉得这个库对您有用",
		FontPath: "../assets/msyh.ttf",
	}
	//绘制文字
	textHandler2 := &TextHandler{
		Next:     Next{},
		X:        180,
		Y:        150,
		Size:     22,
		R:        255,
		G:        241,
		B:        250,
		Text:     "请随意赞赏~~",
		FontPath: "../assets/msyh.ttf",
	}
	//绘制直线
	textLineHandler := &TextLineHandler{
		Next:   Next{},
		X:      175,
		Y:      152,
		Width:  100,
		Height: 10,
		R:      0,
		G:      0,
		B:      0,
		A:      255,
	}
	//绘制矩形
	imageRectangleHandler := &ImageRectangleHandler{
		X:      30,
		Y:      400,
		Width:  100,
		Height: 100,
		R:      0,
		G:      0,
		B:      0,
		A:      255,
	}
	//结束绘制，把前面的内容合并成一张图片
	endHandler := &FileEndHandler{
		Filepath: "../build/poster_" + core.RandString(20) + ".png",
	}

	startHandler := NewNullHandler()

	// 链式调用绘制过程
	startHandler.
		SetNext(backgroundHandler).
		SetNext(imageCircleHandler).
		SetNext(textHandler1).
		SetNext(textHandler2).
		SetNext(textLineHandler).
		SetNext(imageLocalHandler).
		SetNext(qrCodeHandler).
		SetNext(imageRectangleHandler).
		SetNext(endHandler)

	// 开始执行业务
	if err := startHandler.Run(ctx); err != nil {
		// 异常
		fmt.Println("Fail | Error:" + err.Error())
		return
	}
	// 成功
	fmt.Println("Success")
	return
}
