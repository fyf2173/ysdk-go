package kapi

// QueryOrderLogistics 订单物流轨迹
func (kc *Client) QueryOrderLogistics(jdOrderId int64) ([]*LogisticTrace, error) {
	params := new(QueryOrderLogisticsReq)
	params.Protocol = *kc.GetProtocolParams().WithCustomer(kc.CustomerID)
	params.Pin = kc.Pin
	params.OrderId = jdOrderId

	var response []*LogisticTrace
	if err := kc.Request(GetLogistics, params, &response); err != nil {
		return nil, err
	}
	return response, nil
}
