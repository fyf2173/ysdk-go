package jcq

type MessageHandlerFunc func(msg *ResponseMessage) error
type MessageHandlerLog func(topic string, resp ResponseOk)

type TopicFunc struct {
	Name  string
	Topic string
	Size  int
	Do    MessageHandlerFunc
	Log   MessageHandlerLog
}

// CmqConsumeAndAck 消费队列消息，并确认消息
func (jc *Client) CmqConsumeAndAck(topicFn TopicFunc, ack bool) error {
	jc.WithSize(topicFn.Size)
	for {
		// err := json.Unmarshal([]byte(`{"topicName":"568091687201$Default$open_message_ct_order_cancel_E5BF4AF84832FA37E1CE6434B6E673D5","ackIndex":"0|1660122345106","messages":[{"messageId":"2315e74f-a777-421c-a25c-5f1a1e2e2756","messageBody":"{\"channelOrderId\":\"232825476519488\",\"dParentOrderId\":0,\"cancelTime\":\"2022-08-08 19:28:40\",\"orderId\":250435739557,\"cancelStatus\":1,\"customerId\":235680011,\"appKey\":\"E5BF4AF84832FA37E1CE6434B6E673D5\",\"channelId\":25003436,\"rootOrderId\":0,\"timestamp\":1659958120693}","properties":{"BUSINESS_ID":"250435739557","PROPERTY_RETRY_TIMES":"0"}},{"messageId":"0a5af153-5485-40fc-8796-413a69f0f021","messageBody":"{\"channelOrderId\":\"233000054007360\",\"dParentOrderId\":0,\"cancelTime\":\"2022-08-09 18:59:37\",\"orderId\":246342349621,\"cancelStatus\":1,\"customerId\":235680011,\"appKey\":\"E5BF4AF84832FA37E1CE6434B6E673D5\",\"channelId\":25003436,\"rootOrderId\":0,\"timestamp\":1660042777357}","properties":{"BUSINESS_ID":"246342349621","PROPERTY_RETRY_TIMES":"0"}},{"messageId":"29e1acad-43af-4618-bb45-030a36bacdaa","messageBody":"{\"channelOrderId\":\"233000814480960\",\"dParentOrderId\":0,\"cancelTime\":\"2022-08-09 18:59:47\",\"orderId\":250451216802,\"cancelStatus\":1,\"customerId\":235680011,\"appKey\":\"E5BF4AF84832FA37E1CE6434B6E673D5\",\"channelId\":25003436,\"rootOrderId\":0,\"timestamp\":1660042787545}","properties":{"BUSINESS_ID":"250451216802","PROPERTY_RETRY_TIMES":"0"}}]}`), &resp)
		var resp ResponseOk
		if err := jc.Consume(topicFn.Topic, &resp); err != nil {
			return err
		}
		if resp.AckIndex == "" {
			return nil
		}
		if topicFn.Log != nil {
			topicFn.Log(topicFn.Topic, resp)
		}

		var shouldAck = true
		for _, msg := range resp.Messages {
			if err := topicFn.Do(msg); err != nil {
				shouldAck = false
				continue
			}
		}
		if ack && shouldAck {
			if err := jc.Ack(topicFn.Topic, ACK_SUCCESS, resp.AckIndex, nil); err != nil {
				return err
			}
		}
	}
}
