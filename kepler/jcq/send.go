package jcq

import (
	"net/http"
)

// Send 发送消息
func (jc *Client) Send(topic string, msgType string, msgs []*SendMessageRequestMsg, resp interface{}) error {
	jc.WithTopic(topic)
	var params = SendMessageRequest{
		Topic:    jc.Topic,
		MsgType:  msgType,
		Messages: msgs,
	}
	return jc.Request(http.MethodPost, consumerPath, params, resp)
}
