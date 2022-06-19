package wechat

import "encoding/xml"

// SendRedpackReq 发送红包请求
type SendRedpackReq struct {
	XMLName     xml.Name `xml:"root" json:"-"`
	NonceStr    string   `xml:"nonce_str" json:"nonce_str"`
	Sign        string   `xml:"sign" json:"sign"`
	MchBillno   string   `xml:"mch_billno" json:"mch_billno"`
	MchID       string   `xml:"mch_id" json:"mch_id"`
	Wxappid     string   `xml:"wxappid" json:"wxappid"`
	SendName    string   `xml:"send_name" json:"send_name"`
	ReOpenid    string   `xml:"re_openid" json:"re_openid"`
	TotalAmount int      `xml:"total_amount" json:"total_amount"` // 付款金额，单位分
	TotalNum    int      `xml:"total_num" json:"total_num"`       // 红包发放总人数
	Wishing     string   `xml:"wishing" json:"wishing"`           // 红包祝福语
	ClientIP    string   `xml:"client_ip" json:"client_ip"`
	ActName     string   `xml:"act_name" json:"act_name"`
	Remark      string   `xml:"remark" json:"remark"`
	SceneID     string   `xml:"scene_id,omitempty" json:"scene_id,omitempty"` // 发放红包使用场景，红包金额大于200或者小于1元时必传
	RiskInfo    string   `xml:"risk_info,omitempty" json:"risk_info,omitempty"`
}

// SendRedpackResp 发送红包返回
type SendRedpackResp struct {
	WechatResp
	ResultCode  string `xml:"result_code" json:"result_code"`
	ErrCode     string `xml:"err_code" json:"err_code"`
	ErrCodeDes  string `xml:"err_code_des" json:"err_code_des"`
	MchBillno   string `xml:"mch_billno" json:"mch_billno"`
	MchID       string `xml:"mch_id" json:"mch_id"`
	Wxappid     string `xml:"wxappid" json:"wxappid"`
	ReOpenid    string `xml:"re_openid" json:"re_openid"`
	TotalAmount int    `xml:"total_amount" json:"total_amount"`
	SendListid  string `xml:"send_listid" json:"send_listid"` // 红包订单的微信单号
}
