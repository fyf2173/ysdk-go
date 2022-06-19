package wechat

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
)

const (
	UnifiedOrderUrl = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	RefundOrderUrl  = "https://api.mch.weixin.qq.com/secapi/pay/refund"
	QueryOrderUrl   = "https://api.mch.weixin.qq.com/pay/orderquery"
	SendRedpackUrl  = "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack"
	codeSuccess     = "SUCCESS"
	codeFail        = "FAIL"
	RefundSuccess   = "SUCCESS"
	RefundChange    = "CHANGE"
	RefundClose     = "REFUNDCLOSE"
)

var WechatClient *http.Client
var logger *log.Logger

// ComResp 通用返回数据
type WechatResp struct {
	XMLName    xml.Name `xml:"xml" json:"-"`                   // 指定最外层的标签名
	ReturnCode string   `xml:"return_code" json:"return_code"` // 返回状态码
	ReturnMsg  string   `xml:"return_msg" json:"return_msg"`   // 返回信息
}

// SetLogger 设置logger
func SetLogger(newLogger *log.Logger) {
	logger = newLogger
	return
}

// loggerPrint 打印日志
func loggerPrint(data interface{}) {
	if logger != nil {
		b, _ := json.Marshal(data)
		logger.Println(string(b))
	}
}

// MakeSign 签名算法
func MakeSign(data interface{}, skipSign bool, appSecret string) (string, error) {
	var (
		stringA  string
		keySlice []string
		dataMap  = make(map[string]interface{})
	)
	b, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	if err = json.Unmarshal(b, &dataMap); err != nil {
		return "", err
	}
	for k := range dataMap {
		keySlice = append(keySlice, k)
	}
	sort.Strings(keySlice)
	for i := 0; i <= len(keySlice)-1; i++ {
		if dataMap[keySlice[i]] == "" {
			continue
		}
		// 验证调用返回或微信主动通知签名时，传送的sign参数不参与签名，将生成的签名与该sign值作校验
		if skipSign && keySlice[i] == "sign" {
			continue
		}
		tmp := dataMap[keySlice[i]]
		if stringA == "" {
			stringA = fmt.Sprintf("%s=%+v", keySlice[i], tmp)
		} else {
			stringA += fmt.Sprintf("&%s=%+v", keySlice[i], tmp)
		}
	}
	stringA += fmt.Sprintf("&key=%s", appSecret)
	m := md5.New()
	m.Write([]byte(stringA))
	return strings.ToUpper(hex.EncodeToString(m.Sum(nil))), nil
}
