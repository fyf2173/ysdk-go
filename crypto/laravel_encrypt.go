package crypto

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"strings"

	"github.com/forgoer/openssl"
	"github.com/techoner/gophp/serialize"
)

type LaravelToken struct {
	appkey string
}

func NewLaravelToken(appkey string) *LaravelToken {
	return &LaravelToken{appkey: appkey}
}

// 加密
func (lt *LaravelToken) Encrypt(value string) (string, error) {
	iv := make([]byte, 16)
	_, err := rand.Read(iv)
	if err != nil {
		return "", err
	}

	//反序列化
	message, err := serialize.Marshal(value)
	if err != nil {
		return "", err
	}

	key := lt.getKey()

	//加密value
	res, err := openssl.AesCBCEncrypt(message, []byte(key), iv, openssl.PKCS7_PADDING)
	if err != nil {
		return "", err
	}

	//base64加密
	resVal := base64.StdEncoding.EncodeToString(res)
	resIv := base64.StdEncoding.EncodeToString(iv)

	//生成mac值
	data := resIv + resVal
	mac := computeHmacSha256(data, key)

	//构造ticket结构
	ticket := make(map[string]interface{})
	ticket["iv"] = resIv
	ticket["mac"] = mac
	ticket["value"] = resVal

	//json序列化
	resTicket, err := json.Marshal(ticket)
	if err != nil {
		return "", err
	}
	//base64加密ticket
	ticketR := base64.StdEncoding.EncodeToString(resTicket)

	return ticketR, nil
}

// 解密
func (lt *LaravelToken) Decrypt(value string) (string, error) {
	//base64解密
	token, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return "", err
	}

	//json反序列化
	tokenJson := make(map[string]string)
	err = json.Unmarshal(token, &tokenJson)
	if err != nil {
		return "", err
	}

	tokenJsonIv, okIv := tokenJson["iv"]
	tokenJsonValue, okValue := tokenJson["value"]
	tokenJsonMac, okMac := tokenJson["mac"]
	if !okIv || !okValue || !okMac {
		return "", errors.New("value is not full")
	}

	key := lt.getKey()

	//mac检查，防止数据篡改
	data := tokenJsonIv + tokenJsonValue
	check := checkMAC(data, tokenJsonMac, key)
	if !check {
		return "", errors.New("mac valid failed")
	}

	//base64解密iv和value
	tokenIv, err := base64.StdEncoding.DecodeString(tokenJsonIv)
	if err != nil {
		return "", err
	}
	tokenValue, err := base64.StdEncoding.DecodeString(tokenJsonValue)
	if err != nil {
		return "", err
	}
	//aes解密value
	dst, err := openssl.AesCBCDecrypt(tokenValue, []byte(key), tokenIv, openssl.PKCS7_PADDING)
	if err != nil {
		return "", err
	}

	//反序列化
	res, err := serialize.UnMarshal(dst)
	if err != nil {
		return "", err
	}
	return res.(string), nil
}

// 处理密钥
func (lt *LaravelToken) getKey() string {
	appKey := lt.appkey
	if strings.HasPrefix(appKey, "base64:") {
		split := appKey[7:]
		if key, err := base64.StdEncoding.DecodeString(split); err == nil {
			return string(key)
		}
		return split
	}
	return appKey
}

// 比较预期的hash和实际的hash
func checkMAC(message, msgMac, secret string) bool {
	expectedMAC := computeHmacSha256(message, secret)
	return hmac.Equal([]byte(expectedMAC), []byte(msgMac))
}

// 计算mac值
func computeHmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}
