package kapi

// GetSkuList 获取sku列表
func (kc *Client) GetSkuList(req QuerySkuListReq) (*QuerySkuListResp, error) {
	paramsMap := make(map[string]interface{})
	paramsMap["ctpProtocol"] = kc.GetProtocolParams().WithCustomer(kc.CustomerID)
	paramsMap["apiSkuListParam"] = req

	var response QuerySkuListResp
	if err := kc.Request(GetSkuListPath, paramsMap, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetBrotherSkus 获取兄弟sku列表
func (kc *Client) GetBrotherSkus(req QueryBrotherSkuListReq) ([]*QueryBrotherSkuListItem, error) {
	paramsMap := make(map[string]interface{})
	paramsMap["ctpProtocol"] = kc.GetProtocolParams().WithCustomer(kc.CustomerID).WithTraceId()
	paramsMap["apiBrotherListParam"] = req

	var response []*QueryBrotherSkuListItem
	if err := kc.Request(GetSkuListPath, paramsMap, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetSkuDetail 获取sku详情
func (kc *Client) GetSkuDetail(req QuerySkuDetailReq) ([]*QuerySkuDetailResp, error) {
	paramsMap := make(map[string]interface{})
	paramsMap["ctpProtocol"] = kc.GetProtocolParams().WithCustomer(kc.CustomerID)
	paramsMap["skuDetailParam"] = req

	var response []*QuerySkuDetailResp
	if err := kc.Request(GetSkuDetailPtah, paramsMap, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetSkuPrice 批量获取sku价格
func (kc *Client) GetSkuPrice(req QuerySkuPriceReq) (*QuerySkuPriceResp, error) {
	paramsMap := make(map[string]interface{})
	paramsMap["ctpProtocol"] = kc.GetProtocolParams().WithCustomer(kc.CustomerID)
	paramsMap["skuPriceInfoParam"] = req

	var response QuerySkuPriceResp
	if err := kc.Request(GetSkuPriceInfoListPath, paramsMap, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
