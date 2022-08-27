package kapi

// QueryAreaStockState 查询区域库存状态
func (kc *Client) QueryAreaStockState(req StockStateReq) ([]*StockStateRespItem, error) {
	paramsMap := make(map[string]interface{})
	paramsMap["ctpProtocol"] = kc.GetProtocolParams().WithCustomer(kc.CustomerID).WithOpName(kc.Pin)
	paramsMap["stockStateParam"] = req

	var response []*StockStateRespItem
	if err := kc.Request(GetAreaStockState, paramsMap, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// QueryFreightFee 获取运费
func (kc *Client) QueryFreightFee(req FreightFeeReq) (float64, error) {
	paramsMap := make(map[string]interface{})
	paramsMap["protocol"] = kc.GetProtocolParams().WithCustomer(kc.CustomerID)
	req.Pin = kc.Pin
	req.PaymentType = PaymentOnline
	paramsMap["apiFreightFeeParam"] = req

	var response FreightFeeResp
	if err := kc.Request(GetFreightFee, paramsMap, &response); err != nil {
		return 0, err
	}
	return response.FreightFee, nil
}

// QueryShipmentType 配送方式
func (kc *Client) QueryShipmentType(req ShipmentTypeReq) (*ShipmentTypeResp, error) {
	paramsMap := make(map[string]interface{})
	paramsMap["protocol"] = kc.GetProtocolParams().WithCustomer(kc.CustomerID)
	req.Pin = kc.Pin
	req.PaymentType = PaymentOnline
	paramsMap["apiShipmentTypeParam"] = req

	var response ShipmentTypeResp
	if err := kc.Request(GetShipmentType, paramsMap, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
