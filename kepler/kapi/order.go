package kapi

import "fmt"


// SubmitOrder 创建订单
func (kc *Client) SubmitOrder(req SubmitOrderReq) (*SubmitOrderResp, error) {
	paramsMap := make(map[string]interface{})
	paramsMap["protocol"] = kc.GetProtocolParams().WithCustomer(kc.CustomerID)
	req.Pin = kc.Pin
	paramsMap["param"] = req

	var response SubmitOrderResp
	if err := kc.Request(CreateOrder, paramsMap, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// QuerySubmitOrder 反查订单
func (kc *Client) QuerySubmitOrder(channelOrderId int64) (*SubmitOrderResp, error) {
	var params = QuerySubmitOrderReq{}
	params.Protocol = kc.GetProtocolParams().WithCustomer(kc.CustomerID)
	params.ChannelOrderId = fmt.Sprintf("%d", channelOrderId)
	params.Pin = kc.Pin

	var response SubmitOrderResp
	if err := kc.Request(QueryOrder, params, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// QueryOrderDetail 查询订单详情
func (kc *Client) QueryOrderDetail(orderId int64) (*QueryOrderDetailResp, error) {
	var params = QueryOrderDetailReq{}
	params.Protocol = kc.GetProtocolParams().WithCustomer(kc.CustomerID)
	params.OrderId = orderId
	params.Pin = kc.Pin

	var response QueryOrderDetailResp
	if err := kc.Request(GetOrderDetail, params, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ConfirmDelivery 订单确认收货
func (kc *Client) ConfirmDelivery(orderId int64) error {
	var params = ConfirmDeliveryReq{}
	params.Protocol = kc.GetProtocolParams().WithCustomer(kc.CustomerID)
	params.ClientPort = "80"
	params.OrderId = orderId

	if err := kc.Request(ConfirmDelivery, params, nil); err != nil {
		return err
	}
	return nil
}

// PushOrder 推送订单（慎用，需jd运营配置后才生效）
func (kc *Client) PushOrder(orderId int64) error {
	var params = PushOrderReq{}
	params.Protocol = kc.GetProtocolParams().WithCustomer(kc.CustomerID)
	params.OrderId = orderId

	if err := kc.Request(PushOrder, params, nil); err != nil {
		return err
	}
	return nil
}

// QueryOrderPayInfo 查询订单支付信息
func (kc *Client) QueryOrderPayInfo(orderId int64) (*OrderPayInfoResp, error) {
	var params = QueryOrderDetailReq{}
	params.Protocol = kc.GetProtocolParams().WithCustomer(kc.CustomerID)
	params.Pin = kc.Pin
	params.OrderId = orderId

	var response OrderPayInfoResp
	if err := kc.Request(GetOrderPayInfo, params, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// QueryInvoiceDetail 查询发票信息
func (kc *Client) QueryInvoiceDetail(orderId int64) ([]*InvoiceDetail, error) {
	var params = QueryOrderDetailReq{}
	params.Protocol = kc.GetProtocolParams().WithCustomer(kc.CustomerID)
	params.Pin = kc.Pin
	params.OrderId = orderId

	var response []*InvoiceDetail
	if err := kc.Request(GetInvoiceInfo, params, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// CancelOrder 取消订单
func (kc *Client) CancelOrder(req CancelOrderReq) (int, error) {
	var params = CancelOrderParams{}
	params.Protocol = kc.GetProtocolParams().WithCustomer(kc.CustomerID)
	params.Pin = kc.Pin
	params.CancelOrderReq = req

	var response CancelOrderResp
	if err := kc.Request(CancelOrder, params, &response); err != nil {
		return 0, err
	}
	return response.CancelStatus, nil
}
