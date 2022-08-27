package jcq

import (
	"net/http"
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
