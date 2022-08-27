package jcq

const (
	MqCancelOrderSuccess = 1
	MqCancelOrderFailed  = 2
	MqCancelOrdering     = 3 // 申请中，等待处理
	MqCancelOrderBacking = 4 // 拒收，商品退回
)

type SendMessageResponse struct {
	MessageIds []string `json:"messageIds"`
}

type MessageBase struct {
	Timestamp  int64  `json:"timestamp"`
	CustomerId int64  `json:"customerId"`
	ChannelId  int64  `json:"channelId"`
	AppKey     string `json:"appKey"`
}

type OrderMessageBase struct {
	MessageBase
	OrderId        int64  `json:"orderId"`
	ChannelOrderId string `json:"channelOrderId"`
	RootOrderId    int64  `json:"rootOrderId"`
	DParentOrderId int64  `json:"dParentOrderId"`
}

// OrderCreateResponse 创建订单
type OrderCreateResponse struct {
	OrderMessageBase

	OrderModel  int     `json:"orderModel"`
	Yn          int     `json:"yn"`
	CreateTime  string  `json:"createTime"`
	TotalFee    float64 `json:"totalFee"`
	DiscountFee float64 `json:"discountFee"`
	FreightFee  float64 `json:"freightFee"`

	VirtualSkuId      string `json:"virtualSkuId"`
	AutoCancelSeconds string `json:"autoCancelSeconds"`

	SkuList      []*OrderCreateResponseSkuItem    `json:"skuList"`
	AfsOrderInfo *OrderCreateResponseAfsOrderInfo `json:"afsOrderInfo"`
}

type OrderCreateResponseSkuItem struct {
	SkuId     int64   `json:"skuId"`
	SkuType   int     `json:"skuType"`
	SkuName   string  `json:"skuName"`
	Quantity  int     `json:"quantity"`
	ImgUrl    string  `json:"imgUrl"`
	MainSkuId int64   `json:"mainSkuId"`
	Price     float64 `json:"price"`
	Cid1      int     `json:"cid1"`
	Cid2      int     `json:"cid2"`
	Cid3      int     `json:"cid3"`
}

type OrderCreateResponseAfsOrderInfo struct {
	OrderId       int64  `json:"orderId"`
	ParentOrderId int64  `json:"parentOrderId"`
	ServiceId     string `json:"serviceId"`
}

// OrderBalanceNotEnoughResponse 账户余额不足
type OrderBalanceNotEnoughResponse struct {
	MessageBase

	SecondMerchantNum string `json:"secondMerchantNum"`
	CustomerNum       string `json:"customerNum"`
	FirstMerchantNum  string `json:"firstMerchantNum"`
	Msg               string `json:"msg"`
}

// OrderPayResponse 订单支付成功
type OrderPayResponse struct {
	OrderMessageBase

	RealPayFee float64 `json:"realPayFee"`
	PayTime    string  `json:"payTime"`
}

// OrderStockOutResponse 订单出库
type OrderStockOutResponse struct {
	OrderMessageBase

	Packages []*OrderStockOutResponsePackages `json:"packages"`
}

type OrderStockOutResponsePackages struct {
	WaybillCode      string                             `json:"waybillCode"`
	LogisticsName    string                             `json:"logisticsName"`
	LogisticsCode    string                             `json:"logisticsCode"`
	OutLogisticsCode string                             `json:"outLogisticsCode"`
	OutboundTime     string                             `json:"outboundTime"`
	SkuList          []*OrderStockOutResponsePackageSku `json:"skuList"`
}

type OrderStockOutResponsePackageSku struct {
	SkuId    string `json:"skuId"`
	SkuName  string `json:"skuName"`
	Quantity int    `json:"quantity"`
}

// OrderDeliveredResponse 订单妥投
type OrderDeliveredResponse struct {
	OrderMessageBase

	OpeTime string `json:"opeTime"`
}

// OrderFinishedResponse 订单完成
type OrderFinishedResponse struct {
	OrderMessageBase

	FinishedTime string `json:"finishedTime"`
}

// OrderCancelResponse 订单取消
type OrderCancelResponse struct {
	OrderMessageBase

	CreateTime      string                        `json:"createTime"`
	CancelStatus    int                           `json:"cancelStatus" desc:"取消状态 1、取消成功 ；2、取消失败 ；3、申请取消 ；4、申请拒收；"`
	RootOrderStatus int                           `json:"rootOrderStatus" desc:"订单取消前的状态,rootOrderStatus只有1和为空两种情况，1表示京东未代扣，为空表示京东已代扣"`
	ExtendInfo      OrderCancelResponseExtendInfo `json:"extendInfo"`
}

type OrderCancelResponseExtendInfo struct {
	CancelReasonType string `json:"cancelReasonType"`
	Class            string `json:"class"`
}

type tk struct {
	Timestamp int64 `json:"timestamp"`
}

// OrderRefundResponse 订单退款成功
type OrderRefundResponse struct {
	OrderMessageBase
	CustomerId int    `json:"customerId"`
	AppKey     string `json:"appKey"`
	SkuList    []struct {
		SkuName  string `json:"skuName"`
		SkuId    int64  `json:"skuId"`
		Quantity int    `json:"quantity"`
	} `json:"skuList"`
	RefundType       int         `json:"refundType"`    //订单消息类型 1 代表售后退款 2 代表取消退款 3 订单多支付退款 4 余额提现 5 整单二次退款
	RefundTime       string      `json:"refundTime"`    //2022-08-17 11:33:07
	ServiceNumber    interface{} `json:"serviceNumber"` //京东服务单号(仅售后有值)
	RefundId         string      `json:"refundId"`
	RefundFee        float64     `json:"refundFee"` //金额 单位：元
	RefundDetailList []struct {
		RefundFee  float64 `json:"refundFee"`
		PayType    int     `json:"payType"`
		RefundMode int     `json:"refundMode"`
	} `json:"refundDetailList"`
	Timestamp int64 `json:"timestamp"` //1660707187321
}

// CreateAfsResponse 售后服务单创建成功
type CreateAfsResponse struct {
	MessageBase
	ChannelOrderId    string `json:"channelOrderId"`
	OrderId           int64  `json:"orderId"`
	AfsServiceId      int64  `json:"afsServiceId"`
	AfsApplyId        int64  `json:"afsApplyId"`
	AfsApplyTime      int64  `json:"afsApplyTime"` // 1660560206000
	AfsType           int    `json:"afsType"`
	SkuName           string `json:"skuName"`
	ChannelAfsApplyId string `json:"channelAfsApplyId"`
	SkuId             int64  `json:"skuId"`
}

// AfaStepResultResponse 售后服务单全流程
type AfaStepResultResponse struct {
	MessageBase

	AfsServiceId  int64  `json:"afsServiceId"`
	AfsResultType string `json:"afsResultType"`
	AfsType       int    `json:"afsType"`
	SkuId         int64  `json:"skuId"`
	SkuName       string `json:"skuName"`
	StepType      string `json:"stepType"`
	OrderId       int64  `json:"orderId"`
	OperationDate string `json:"operationDate"`
}

type SkuChange struct {
	MessageBase
	SkuId      int64  `json:"skuId" desc:"京东商品编号"`
	OuterSkuId string `json:"outerSkuId" desc:"渠道商品编号"`
	Type       int64  `json:"type" desc:"0-删除（京东商品删除） 1-新增 2-修改 3-删除恢复（京东商品删除恢复）"`
}

type SkuPriceChange struct {
	MessageBase
	SkuId      int64  `json:"skuId" desc:"京东商品编号"`
	OuterSkuId string `json:"outerSkuId" desc:"渠道商品编号"`
	Type       int64  `json:"type" desc:"2-修改"`
}

type AddressChange struct {
	MessageBase
	ParentId    int64  `json:"parentId" desc:"父区域Id"`
	AreaName    string `json:"areaName" desc:"渠道商品编号"`
	OperateType int64  `json:"operateType" desc:"操作类型:1（插入）2（更新）3（删除)"`
	AreaId      int64  `json:"areaId" desc:"区域id"`
	AreaLevel   int64  `json:"areaLevel" desc:"行政级别"`
}
