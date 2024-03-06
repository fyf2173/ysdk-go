package util

import (
	"bytes"
	"fmt"
	"io"
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
	return viper.UnmarshalKey(key, node, func(m *mapstructure.DecoderConfig) { m.TagName = "yaml" })
}

func ViperMustGetNode[T any](key string, node T) T {
	if err := viper.UnmarshalKey(key, &node, func(m *mapstructure.DecoderConfig) { m.TagName = "yaml" }); err != nil {
		panic(err)
	}
	return node
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

// GetRandString n is the len of returned rand string
func GetRandString(n int) string {
	var rands = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = rands[rand.Intn(len(rands))]
	}
	return string(b)
}

// GetRandBetween 生成区间随机数，最小是min，最大是max
func GetRandBetween(min, max int64) int64 {
	if min > max || min == 0 || max == 0 {
		return max
	}
	max = max + 1
	return rand.Int63n(max-min) + min
}

const (
	GBK     string = "GBK"
	UTF8    string = "UTF8"
	UNKNOWN string = "UNKNOWN"
)

// 需要说明的是，isGBK()是通过双字节是否落在gbk的编码范围内实现的，
// 而utf-8编码格式的每个字节都是落在gbk的编码范围内，
// 所以只有先调用isUtf8()先判断不是utf-8编码，再调用isGBK()才有意义
func GetStrCoding(data []byte) string {
	if isUTF8(data) == true {
		return UTF8
	} else if isGBK(data) == true {
		return GBK
	} else {
		return UNKNOWN
	}
}

// isGBK check data encode is gbk
func isGBK(data []byte) bool {
	bn := len(data)
	for i := 0; i <= bn-1; i++ {
		single := data[i]
		if single <= 0x7f { // 编码0~127,只有一个字节的编码，兼容ASCII码
			continue
		}
		// 大于127的使用双字节编码，落在gbk编码范围内的字符
		i += 1
		if single >= 0x81 && single <= 0xFE {
			if i >= bn {
				return false
			}
			double := data[i]
			if double >= 0x40 && double <= 0xFE && double != 0x7F {
				continue
			}
		}
	}
	return true
}

func preNUm(data byte) int {
	var mask byte = 0x80
	var num int = 0
	//8bit中首个0bit前有多少个1bits
	for i := 0; i < 8; i++ {
		if (data & mask) != mask {
			break
		}
		num++
		mask = mask >> 1
	}
	return num
}

func isUTF8(data []byte) bool {
	for i := 0; i <= len(data)-1; i++ {
		// 0XXX_XXXX
		if (data[i] & 0x80) == 0x00 {
			continue
		}

		// preNUm() 返回首个字节的8个bits中首个0bit前面1bit的个数，该数量也是该字符所使用的字节数
		num := preNUm(data[i])
		if num <= 2 {
			return false
		}

		// 110X_XXXX 10XX_XXXX
		// 1110_XXXX 10XX_XXXX 10XX_XXXX
		// 1111_0XXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
		// 1111_10XX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
		// 1111_110X 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
		for j := 0; j < num-1; j++ {
			//判断后面的 num - 1 个字节是不是都是10开头
			if (data[i+1] & 0xc0) != 0x80 {
				return false
			}
			i += 1
		}
	}
	return true
}

// GbkToUtf8 transform GBK bytes to UTF-8 bytes
func GbkToUtf8(str []byte) (b []byte, err error) {
	r := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GB18030.NewDecoder())
	b, err = io.ReadAll(r)
	if err != nil {
		return
	}
	return
}

// Utf8ToGbk transform UTF-8 bytes to GBK bytes
func Utf8ToGbk(str []byte) (b []byte, err error) {
	r := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GB18030.NewEncoder())
	b, err = io.ReadAll(r)
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

// StrTransform transform GBK string to UTF-8 string，if str is UTF-8 string，return itself，then transform to UTF-8 string and returned，
// otherwise return origin str
func StrTransform(str string) string {
	if encode := GetStrCoding([]byte(str)); encode == UTF8 {
		return str
	}
	newByte, err := GbkToUtf8([]byte(str))
	if err != nil {
		return str
	}
	return string(newByte)
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

// Stringf 字符格式拼接
func Stringf(args ...string) string {
	var buf strings.Builder
	for _, s := range args {
		buf.WriteString(s)
	}
	return buf.String()
}

// GetProbability 按百分比计算概率，probSeeds为概率种子数组，例：[]int{1, 1, 1, 1, 0, 0, 0, 0, 0, 0}，取概率值为40%
func GetProbability(probSeeds [10]int) bool {
	if index := GetRandBetween(1, int64(len(probSeeds))); probSeeds[index-1] <= 0 {
		return false
	}
	return true
}

// NewProbSeeds 初始化概率种子，probability为[0,1]之间的概率
func NewProbSeeds(probability float64) [10]int {
	var rateArr [10]int
	for i := 0; i < int(probability*10); i++ {
		if i >= len(rateArr) {
			break
		}
		rateArr[i] = 1
	}
	return rateArr
}

// SliceCompareGroup compare two slices and return sames, ldiff, rdiff. sames is the same elements between ref and
// other, ldiff is the diffrence of ref to sames, rdiff is the diffrence of other to sames
func SliceCompareGroup[T comparable](refs, others []T) (sames, ldiff, rdiff []T) {
	for _, rv := range refs {
		if SliceFindOk(rv, others) {
			sames = append(sames, rv)
		}
	}
	for _, rv := range refs {
		if !SliceFindOk(rv, sames) {
			ldiff = append(ldiff, rv)
		}
	}
	for _, ov := range others {
		if !SliceFindOk(ov, sames) {
			rdiff = append(rdiff, ov)
		}
	}
	return
}

// SliceFindOk find val from cmps, if found return true, or return false
func SliceFindOk[T comparable](val T, cmps []T) bool {
	return SliceFindFilterOk(val, cmps, func(val, cmpsval T) bool {
		return val == cmpsval
	})
}

// SliceFindFilter find val from cmps by filter, if found return true, or return false
func SliceFindFilterOk[T comparable](val T, cmps []T, fn func(val, cmpsval T) bool) bool {
	for _, v := range cmps {
		if fn(val, v) == true {
			return true
		}
	}
	return false
}

// SliceFieldFiltered take val slice from refs slice by filtered
func SliceFieldFilteredSlice[T any, R any](refs []T, fn func(val T) (R, error)) []R {
	var result []R
	for _, v := range refs {
		tmp, err := fn(v)
		if err != nil {
			continue
		}
		result = append(result, tmp)
	}
	return result
}

// SliceFieldFilteredMapWithKey take field map with customed key from refs slice by filtered
func SliceFieldFilteredMapWithKey[T any, K, R comparable](refs []T, fn func(val T) (K, R, error)) map[K]R {
	var result = make(map[K]R)
	for _, v := range refs {
		key, tmp, err := fn(v)
		if err != nil {
			continue
		}
		result[key] = tmp
	}
	return result
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MapFieldFilteredSlice take field slice from refs map by filtered
func MapFieldFilteredSlice[K comparable, R any, T any](refs map[K]R, fn func(key K, val R) (T, error)) []T {
	var result []T
	for k, r := range refs {
		tmp, err := fn(k, r)
		if err != nil {
			continue
		}
		result = append(result, tmp)
	}
	return result
}
