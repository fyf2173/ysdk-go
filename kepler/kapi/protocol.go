package kapi

type Protocol struct {
	CustomerId int64  `json:"customerId"`
	ChannelId  int64  `json:"channelId"`
	TraceId    string `json:"traceId,omitempty"`
	AppKey     string `json:"appKey"`
	ClientIp   string `json:"clientIp,omitempty"`
	OpName     string `json:"OpName,omitempty"`
	ParentId   int64  `json:"parentId,omitempty"`
}

func (ptl *Protocol) WithCustomer(customerId int64) *Protocol {
	ptl.CustomerId = customerId
	return ptl
}

func (ptl *Protocol) WithClientIP(clientIp string) *Protocol {
	ptl.ClientIp = clientIp
	return ptl
}

func (ptl *Protocol) WithOpName(opName string) *Protocol {
	ptl.OpName = opName
	return ptl
}
