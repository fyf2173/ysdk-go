package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// TrimFloatZeroSuffix 格式化浮点数，去除小数点末尾的0或.
func TrimFloatZeroSuffix(floatStr string) string {
	dotIndex := strings.Index(floatStr, ".")
	if dotIndex < 0 {
		return floatStr
	}
	breakIndex := 0
	for i := len(floatStr) - 1; i >= dotIndex; i-- {
		tmp := string(floatStr[i])
		if tmp != "0" {
			if tmp == "." {
				i--
			}
			breakIndex = i
			break
		}
	}
	return floatStr[:breakIndex+1]
}

// Assert 断言
func Assert(err error) {
	if err != nil {
		panic(err)
	}
}

func MultiAssert(errs ...error) {
	for _, err := range errs {
		if err != nil {
			panic(err)
		}
	}
}

func MultiCheck(errs ...error) error {
	for i := 0; i <= len(errs)-1; i++ {
		if errs[i] != nil {
			return errs[i]
		}
	}
	return nil
}

func PickUnusedPort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	port := l.Addr().(*net.TCPAddr).Port
	if err := l.Close(); err != nil {
		return 0, err
	}
	return port, nil
}

func NanotimeToDatetime(nanoTime int64) string {
	return time.Unix(nanoTime/1e9, 0).Format("2006-01-02 15:04:05")
}

func TimeFixDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func ViperGetNode(key string, node interface{}) error {
	return viper.UnmarshalKey(key, node, func(m *mapstructure.DecoderConfig) {
		m.TagName = "yaml"
	})
}

// Validate 校验参数
func Validate(params interface{}) error {
	reflectVal := reflect.ValueOf(params)
	if reflectVal.Type().Kind() == reflect.Ptr {
		reflectVal = reflectVal.Elem()
	}

	if err := validator.New().Struct(params); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, v := range validationErrors {
			structField, _ := reflectVal.Type().FieldByName(v.Field())
			if msg := structField.Tag.Get(v.Tag()); msg != "" {
				return fmt.Errorf(msg)
			}
		}
		return err
	}
	return nil
}

// n is the len of returned rand string
func GetRandString(n int) string {
	var rands = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = rands[rand.Intn(len(rands))]
	}
	return string(b)
}

// GetRandBetween 生成区间随机数
func GetRandBetween(min, max int64) int64 {
	if min > max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

// GbkToUtf8 transform GBK bytes to UTF-8 bytes
func GbkToUtf8(str []byte) (b []byte, err error) {
	r := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GB18030.NewDecoder())
	b, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}
	return
}

// Utf8ToGbk transform UTF-8 bytes to GBK bytes
func Utf8ToGbk(str []byte) (b []byte, err error) {
	r := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GB18030.NewEncoder())
	b, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}
	return
}

// StrToUtf8 transform GBK string to UTF-8 string and replace it, if transformed success, returned nil error, or died by error message
func StrToUtf8(str *string) error {
	b, err := GbkToUtf8([]byte(*str))
	if err != nil {
		return err
	}
	*str = string(b)
	return nil
}

// StrToGBK transform UTF-8 string to GBK string and replace it, if transformed success, returned nil error, or died by error message
func StrToGBK(str *string) error {
	b, err := Utf8ToGbk([]byte(*str))
	if err != nil {
		return err
	}
	*str = string(b)
	return nil
}

// Substr returns the substr from start to length, if length smaller than 0, Substr returns the substr from start to end
func Substr(s string, start, length int) string {
	bt := []rune(s)
	if len(bt) <= 0 {
		return ""
	}
	if start < 0 {
		start = 0
	}
	if start > len(bt) {
		start = start % len(bt)
	}
	var end int
	if length <= 0 {
		return string(bt[start:])
	}
	if (start + length) > (len(bt) - 1) {
		end = len(bt)
	} else {
		end = start + length
	}
	return string(bt[start:end])
}

func StrStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if !strings.Contains(haystack, needle) {
		return -1
	}
	return strings.Index(haystack, needle)
}

func CallFuncs(fc interface{}, params ...interface{}) (result interface{}, err error) {
	f := reflect.ValueOf(fc)
	if f.Kind() != reflect.Func {
		err = fmt.Errorf("fc is not func")
		return
	}
	if len(params) != f.Type().NumIn() {
		err = fmt.Errorf("the number of params is not adapted")
		return
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	resp := f.Call(in)
	if len(resp) > 0 {
		result = reflect.ValueOf(resp[0].Interface()).Interface()
		return
	}
	result = nil
	return
}
