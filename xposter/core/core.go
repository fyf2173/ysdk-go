/**
 * @Author: entere@126.com
 * @Description:
 * @File:  core
 * @Version: 1.0.0
 * @Date: 2020/4/22 10:41
 */

package core

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
)

// Rect 新PNG载体
type Rect struct {
	X0 int
	X1 int
	Y0 int
	Y1 int
}

// Pt 坐标
type Pt struct {
	X int
	Y int
}

// DImage 图片切片
type DImage struct {
	PNG draw.Image //合并到的PNG切片,可用image.NewrRGBA设置
	X   int        //横坐标
	Y   int        //纵坐标
}

// DText 文字切片
type DText struct {
	PNG   draw.Image //合并到的PNG切片,可用image.NewrRGBA设置
	Title string     //文字
	X     int        //横坐标
	Y     int        //纵坐标
	Size  float64
	R     uint8
	G     uint8
	B     uint8
	A     uint8
}

// Rectangle 设置矩形
type Rectangle struct {
	R      uint8
	G      uint8
	B      uint8
	A      uint8
	Width  int
	Height int
}

// NewMerged 新建文件载体
func NewMerged(path string) (*os.File, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// NewPNG 新建图片载体
func NewPNG(X0 int, Y0 int, X1 int, Y1 int) *image.RGBA {
	return image.NewRGBA(image.Rect(X0, Y0, X1, Y1))
}

// MergeImage 合并图片到载体
func MergeImage(PNG draw.Image, image image.Image, imageBound image.Point) {
	draw.Draw(PNG, PNG.Bounds(), image, imageBound, draw.Over)
}

// LoadTextType 读取字体类型
func LoadTextType(path string) (*truetype.Font, error) {
	fbyte, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	trueTypeFont, err := freetype.ParseFont(fbyte)
	if err != nil {
		return nil, err
	}
	return trueTypeFont, nil
}

// NewDrawText 创建新字体切片
func NewDrawText(png draw.Image) *DText {
	return &DText{
		PNG:  png,
		Size: 18,
		X:    0,
		Y:    0,
		R:    0,
		G:    0,
		B:    0,
		A:    255,
	}
}

// SetColor 设置字体颜色
func (dtext *DText) SetColor(R uint8, G uint8, B uint8) {
	dtext.R = R
	dtext.G = G
	dtext.B = B
}

// MergeText 合并字体到载体
func (dtext *DText) MergeText(title string, size float64, tf *truetype.Font, x int, y int) error {
	fc := freetype.NewContext()
	//设置屏幕每英寸的分辨率
	fc.SetDPI(72)
	//设置用于绘制文本的字体
	fc.SetFont(tf)
	//以磅为单位设置字体大小
	fc.SetFontSize(size)
	//设置剪裁矩形以进行绘制
	fc.SetClip(dtext.PNG.Bounds())
	//设置目标图像
	fc.SetDst(dtext.PNG)
	//设置绘制操作的源图像，通常为 image.Uniform
	fc.SetSrc(image.NewUniform(color.RGBA{dtext.R, dtext.G, dtext.B, dtext.A}))

	pt := freetype.Pt(x, y)
	_, err := fc.DrawString(title, pt)
	if err != nil {
		return err
	}
	return nil
}

// Merge 合并到图片
func Merge(png draw.Image, merged *os.File) error {
	err := jpeg.Encode(merged, png, nil)
	if err != nil {
		return err
	}
	return nil
}

// DrawQRImage 获取二维码图像
func DrawQRImage(url string, level qrcode.RecoveryLevel, size int) (image.Image, error) {
	newQr, err := qrcode.New(url, level)
	if err != nil {
		return nil, err
	}
	qrImage := newQr.Image(size)
	return qrImage, nil
}

// GetResourceReader 获取资源
func GetResourceReader(url string, width int, hight int) (newImg *image.RGBA, err error) {
	r := bytes.NewReader(nil)
	if url[0:4] == "http" {
		resp, err := http.Get(url)
		if err != nil {
			return newImg, err
		}
		defer resp.Body.Close()
		fileBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return newImg, err
		}
		r = bytes.NewReader(fileBytes)
	} else {
		fileBytes, err := ioutil.ReadFile(url)
		if err != nil {
			return newImg, err
		}
		r = bytes.NewReader(fileBytes)
	}
	img, _, err := image.Decode(r)

	// 调用resize库进行图片缩放(高度填0，resize.Resize函数中会自动计算缩放图片的宽高比)
	m := resize.Resize(uint(width), uint(hight), img, resize.Lanczos3)

	newImg = NewPNG(0, 0, width, hight)                              //创建一个新RGBA图像
	draw.Draw(newImg, newImg.Bounds(), m, m.Bounds().Min, draw.Over) //画上缩放后的图片

	return newImg, nil
}

// RandString 生成随机字符串
func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

// ImageRectangle 绘制矩形
func ImageRectangle(req Rectangle) (*image.RGBA, error) {
	m := NewPNG(0, 0, req.Width, req.Height)
	blue := color.RGBA{req.R, req.G, req.B, req.A}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	return m, nil
}
