package jcq

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type CommonConsumerResp struct {
	RequestId string                 `json:"requestId"`
	Result    json.RawMessage        `json:"result"`
	Error     *CommonConsumerRespErr `json:"error,omitempty"`
}

type CommonConsumerRespErr struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type ConsumerReq struct {
	Topic                string `json:"topic"`
	ConsumerGroupId      string `json:"consumerGroupId"`
	Size                 int    `json:"size,omitempty" desc:"一次最多拉取消息条数，0 < size <=32，defaultValue = 32"`
	ConsumerId           int64  `json:"consumerId,omitempty" desc:"defaultValue = httpProxyId"`
	ConsumeFromWhere     string `json:"consumeFromWhere,omitempty" desc:"默认的起始消费位置，可选值：HEAD、TAIL，defaultValue = HEAD"`
	FilterExpressionType string `json:"filterExpressionType,omitempty" desc:"消息过滤表达式类型，目前可选值只有TAG"`
	FilterExpression     string `json:"filterExpression,omitempty" desc:"消息过滤表达式，默认没有过滤，如果需要过滤，此参数与filterExpressionType需同时传入"`
	Ack                  string `json:"ack,omitempty" desc:"拉消息时是否由服务端自动ACK，ACK 必须在消息确认接收的超时时间内，可选值true、false，为true时，服务端自动ACK消费的消息，为false时，需要客户端来ACK消费的消息，默认值defaultValue = false"`
}

// Consume 消费消息
func (jc *Client) Consume(topic string, resp interface{}) error {
	jc.WithTopic(topic)
	var params = ConsumerReq{
		Topic:           jc.Topic,
		ConsumerGroupId: jc.ConsumerGroupId,
		Size:            jc.Size,
	}

	vals := url.Values{}
	vals.Set("topic", params.Topic)
	vals.Set("consumerGroupId", params.ConsumerGroupId)
	vals.Set("size", fmt.Sprintf("%d", jc.Size))

	return jc.Request(http.MethodGet, consumerPath+"?"+vals.Encode(), params, resp)
}
