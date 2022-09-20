package kapi

// QueryCanApplyInfo 查询是否可售后详情
func (kc *Client) QueryCanApplyInfo(req CanApplyInfoReq) (*CanApplyInfoItem, error) {
	paramsMap := make(map[string]interface{})
	paramsMap["ctpProtocol"] = kc.GetProtocolParams()
	req.Pin = kc.Pin
	paramsMap["canApplyInfoParam"] = req

	var response *CanApplyInfoItem
	if err := kc.Request(AfsCanApplyInfo, paramsMap, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// QueryApplyReasons 获取申请售后原因列表
func (kc *Client) QueryApplyReasons(req QueryApplyReasonsReq) ([]*ApplyReasonItem, error) {
	paramsMap := make(map[string]interface{})
	paramsMap["ctpProtocol"] = kc.GetProtocolParams()
	req.Pin = kc.Pin
	paramsMap["canApplyInfoParam"] = req

	var response []*ApplyReasonItem
	if err := kc.Request(AfsGetApplyReason, paramsMap, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// CreateAfsApply 申请售后服务单
func (kc *Client) CreateAfsApply(req CreateAfsApplyReq) (*CreateAfsApplyResp, error) {
	paramsMap := make(map[string]interface{})
	paramsMap["ctpProtocol"] = kc.GetProtocolParams()
	req.Pin = kc.Pin
	paramsMap["afsApplyParam"] = req

	var response *CreateAfsApplyResp
	if err := kc.Request(AfsApplyCreate, paramsMap, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetAfsLogisticAddress 获取售后回寄地址
func (kc *Client) GetAfsLogisticAddress(afsServiceId int64) (*GetAfsLogisticAddressResp, error) {
	paramsMap := make(map[string]interface{})
	paramsMap["ctpProtocol"] = kc.GetProtocolParams()
	paramsMap["logisticsAddressParam"] = map[string]interface{}{
		"afsServiceId": afsServiceId,
		"pin":          kc.Pin,
	}

	var response *GetAfsLogisticAddressResp
	if err := kc.Request(AfsApplyCreate, paramsMap, &response); err != nil {
		return nil, err
	}
	return response, nil
}
