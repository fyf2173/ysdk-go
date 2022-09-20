package jcq

import (
	"net/http"
)

const (
	ACK_SUCCESS        = "SUCCESS"        //消费成功
	ACK_CONSUME_FAILED = "CONSUME_FAILED" //消费失败,服务端会进行重新推送
	ACK_RESEND         = "RESEND"         //立即重发
	ACK_DISCARD        = "DISCARD"        //丢弃消息，服务端不会进行重试
)

type MessageAckRequest struct {
	Topic           string `json:"topic"`
	ConsumerGroupId string `json:"consumerGroupId"`
	AckAction       string `json:"ackAction"`
	AckIndex        string `json:"ackIndex"`
}

// Ack 确认消息
func (jc *Client) Ack(topic, ackAction, ackIndex string, resp interface{}) error {
	jc.WithTopic(topic)
	var params = MessageAckRequest{
		Topic:           jc.Topic,
		ConsumerGroupId: jc.ConsumerGroupId,
		AckAction:       ackAction,
		AckIndex:        ackIndex,
	}
	return jc.Request(http.MethodPost, ackPath, params, resp)
}
