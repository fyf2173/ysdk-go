package kapi

// GetAddress 获取京东四级地址
func (kc *Client) GetAddress(parentId int64) ([]AddressData, error) {
	params := kc.GetProtocolParams().WithCustomer(kc.CustomerID)
	params.ParentId = parentId
	var response []AddressData
	if err := kc.Request(GetJdChildAddr, params, &response); err != nil {
		return nil, err
	}
	return response, nil
}
