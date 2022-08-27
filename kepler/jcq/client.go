package jcq

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/fyf2173/ysdk-go/crypto"
)

const (
	//endPoint    = "https://jcq-shared-004.cn-north-1.jdcloud.com"
	endPoint = "https://116.196.69.52:443"
)

const (
	consumerPath = "/v2/messages"
	ackPath      = "/v2/ack"
)

type Header struct {
	AccessKey string `json:"accessKey"`
	Signature string `json:"signature"`
	DateTime  string `json:"dateTime" desc:"ISO-8601:2004"`
}

type Response struct {
	RequestId string      `json:"requestId"`
	Result    ResponseOk  `json:"result"`
	Error     ResponseErr `json:"error"`
}

type ResponseErr struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type ResponseMessage struct {
	MessageId   string `json:"messageId"`
	MessageBody string `json:"messageBody"`
	Properties  struct {
		TAGS string `json:"TAGS"`
	} `json:"properties"`
}

type ResponseOk struct {
	TopicName string `json:"topicName"`
	AckIndex  string `json:"ackIndex"`
	Messages  []*ResponseMessage
}

type SendMessageRequest struct {
	Topic    string                   `json:"topic"`
	MsgType  string                   `json:"type" desc:"NORMAL, ORDER"`
	Messages []*SendMessageRequestMsg `json:"messages"`
}

type SendMessageRequestMsg struct {
	Body         string            `json:"body"`
	DelaySeconds int               `json:"delaySeconds,omitempty"`
	Tag          string            `json:"tag,omitempty"`
	Properties   map[string]string `json:"properties,omitempty"`
}

type Client struct {
	AppKey          string
	TenantId        int64
	AccessKeyId     string
	AccessKeySecret string
	Topic           string
	ConsumerGroupId string
	Size            int
}

func NewClient(appKey string, tenantId int64, accessKey, accessSecret string) *Client {
	return &Client{
		AppKey:          appKey,
		TenantId:        tenantId,
		AccessKeyId:     accessKey,
		AccessKeySecret: accessSecret,
		Size:            1,
	}
}

func (jc *Client) WithTopic(topic string) *Client {
	if jc.TenantId == 0 {
		jc.Topic = fmt.Sprintf("open_message_%s_%s", topic, jc.AppKey)
		return jc
	}
	jc.Topic = fmt.Sprintf("%d$%s$open_message_%s_%s", jc.TenantId, "Default", topic, jc.AppKey)
	return jc
}

func (jc *Client) WithGroupId(groupId string) *Client {
	jc.ConsumerGroupId = groupId
	return jc
}

func (jc *Client) WithSize(size int) *Client {
	if size <= 0 {
		jc.Size = 1
	} else if size >= 32 {
		jc.Size = 32
	} else {
		jc.Size = size
	}
	return jc
}

// GetSignature 加密数据源，生成签名数据
func (jc *Client) GetSignature(signSource string) string {
	mac := hmac.New(sha1.New, []byte(jc.AccessKeySecret))
	mac.Write([]byte(signSource))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func (jc *Client) NewHeader() *Header {
	jh := &Header{}
	jh.AccessKey = jc.AccessKeyId
	jh.DateTime = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	return jh
}

func (jh *Header) GetSendMessageSignSource(msg *SendMessageRequestMsg) string {
	var msgMap = make(map[string]interface{})
	msgMap["body"] = msg.Body
	if msg.DelaySeconds > 0 {
		msgMap["delaySeconds"] = msg.DelaySeconds
	}
	if msg.Tag != "" {
		msgMap["tag"] = msg.Tag
	}
	for k, v := range msg.Properties {
		msgMap[k] = v
	}

	var signKey []string
	for k := range msgMap {
		signKey = append(signKey, k)
	}
	sort.Strings(signKey)

	var signSource string
	for i := 0; i <= len(signKey)-1; i++ {
		signSource += fmt.Sprintf("%s=%+v", signKey[i], msgMap[signKey[i]]) + "&"
	}
	return crypto.Md5Str(strings.TrimRight(signSource, "&"))
}

// GetSignSource 获得签名数据源
func (jh *Header) GetSignSource(params interface{}) string {
	var signMap = make(map[string]interface{})
	signMap["accessKey"] = jh.AccessKey
	signMap["dateTime"] = jh.DateTime
	b, _ := json.Marshal(params)
	_ = json.Unmarshal(b, &signMap)

	if _, ok := signMap["messages"]; ok {
		if sendReq, ok := params.(SendMessageRequest); ok {
			var msgSigns []string
			for _, v := range sendReq.Messages {
				msgSigns = append(msgSigns, jh.GetSendMessageSignSource(v))
			}
			signMap["messages"] = strings.Join(msgSigns, ",")
		}
	}

	var signKey []string
	for k := range signMap {
		signKey = append(signKey, k)
	}
	sort.Strings(signKey)

	var vals string
	for i := 0; i <= len(signKey)-1; i++ {
		vals += fmt.Sprintf("%s=%+v", signKey[i], signMap[signKey[i]]) + "&"
	}
	return strings.TrimRight(vals, "&")
}
