package kapi

const (
	RouterJson              = "https://api.jd.com/routerjson"                         // 正式地址
	GetJdChildAddr          = "jingdong.ctp.order.getChildAreaList"                   // 获取京标四级地址API
	GetAreaStockState       = "jingdong.ctp.ware.stock.queryAreaStockState"           // 库存状态
	GetFreightFee           = "jingdong.ctp.order.getFreightFee"                      // 获取邮费
	GetShipmentType         = "jingdong.ctp.order.getShipmentType"                    // 配送方式
	CreateOrder             = "jingdong.ctp.order.submitOrder"                        // 创建订单
	QueryOrder              = "jingdong.ctp.order.querySubmitOrder"                   // 反查订单
	GetOrderDetail          = "jingdong.ctp.order.getOrderDetail"                     // 订单详情
	GetLogistics            = "jingdong.ctp.order.getLogistics"                       // 订单物流
	ConfirmDelivery         = "jingdong.ctp.order.confirmDelivery"                    // 确认收货
	PushOrder               = "jingdong.ctp.order.pushOrder"                          // 订单推送
	PayOrder                = "jingdong.ctp.order.pushOrder"                          // 订单支付
	GetOrderPayInfo         = "jingdong.ctp.order.getOrderPayInfo"                    // 支付信息
	GetInvoiceInfo          = "jingdong.ctp.finance.getInvoiceDetail"                 // 发票信息
	CancelOrder             = "jingdong.ctp.order.cancelOrder"                        // 取消订单
	AfsCanApplyInfo         = "jingdong.ctp.afs.operate.apply.getIsCanApplyInfo"      // 可否售后
	AfsGetApplyReason       = "jingdong.ctp.afs.operate.apply.getApplyReason"         // 售后原因
	AfsApplyCreate          = "jingdong.ctp.afs.operate.apply.createAfsApply"         // 申请售后
	AfsLogisticAddress      = "jingdong.ctp.afs.logistics.getLogisticsAddress"        // 寄回地址
	AfsLogisticInfo         = "jingdong.ctp.afs.logistics.postBackLogisticsBillParam" // 售后物流
	AfsServiceDetail        = "jingdong.ctp.afs.servicenbill.getAfsServiceDetail"     // 售后详情
	AfsServiceCancel        = "jingdong.ctp.afs.servicenbill.cancelAfsService"        // 取消售后
	GetSkuListPath          = "jingdong.ctp.ware.sku.getSkuList"                      //获取渠道商品列表
	GetSkuDetailPtah        = "jingdong.ctp.ware.sku.getSkuDetail"                    //获取商品详情信息
	GetSkuPriceInfoListPath = "jingdong.ctp.ware.price.getSkuPriceInfoList"           //批量获取sku价格接口
	GetSkuBrotherListPath   = "jingdong.ctp.ware.sku.getBrotherList"                  //获取兄弟SKU列表
)

const (
	GetSkuListTag          = "jingdong_ctp_ware_sku_getSkuList_responce"
	GetSkuDetailTag        = "jingdong_ctp_ware_sku_getSkuDetail_responce"
	GetSkuPriceInfoListTag = "jingdong_ctp_ware_price_getSkuPriceInfoList_responce"
	GetSkuBrotherListTag   = "jingdong_ctp_ware_sku_getBrotherList_responce"
)

const (
	MqSkuChange        = "ct_sku_change"               // 商品信息变更
	MqSkuPriceChange   = "ct_sku_price_change"         // 商品价格变更
	MqAddrChange       = "jd_address_change"           // 四级地址变更
	MqOrderCreate      = "ct_order_create"             // 订单创建成功
	MqOrderPay         = "ct_order_pay"                // 订单支付成功
	MqStockOut         = "ct_order_stockout"           // 订单出库
	MqOrderDelivered   = "ct_order_delivered"          // 订单妥投
	MqOrderFinish      = "ct_order_finish"             // 订单完成
	MqOrderCancel      = "ct_order_cancel"             // 订单取消
	MqOrderRefund      = "ct_order_refund"             // 订单退款成功
	MqAfsCreate        = "ct_afs_create"               // 创建售后
	MqAfsStep          = "ct_afs_step_result"          // 售后单全流程
	MqBalanceNotEnough = "ct_order_balance_not_enough" // 余额不足提醒
)

const (
	CancelCodeForDeliveringDelay        = 1   // 订单不能按预计时间送达
	CancelCodeForLowerPriceOtherChannel = 2   // 其他渠道价格更低
	CancelCodeForLowerPricePlatform     = 3   // 该商品降价了
	CancelCodeForNotWant                = 4   // 不想买了
	CancelCodeForOtherReason            = 5   // 其他原因
	CancelCodeForError                  = 6   // 操作有误
	CancelCodeForOutStock               = 7   // 商品无货
	CancelCodeForOther                  = 100 // 其他
)

const (
	CancelNotPay  = 1 // 未支付取消
	CancelForUser = 2 // 用户取消
	CancelForRisk = 3 // 风控取消
)

const (
	CancelTypeForManual = 1 // 订单取消
	CancelTypeForCutOff = 2 // 订单拦截
)

const (
	PickWareNormal       = 4  // 上门取件
	PickWareMajorNormal  = 8  // 大家电上门取件
	PickWareCustomerSend = 40 // 客户发货
)

const (
	StockStateZero     = 0 // 无货
	StockStateNormal   = 1 // 有货（现货）
	StockStatePurchase = 2 // 采购中
)

const (
	ShipmentTypeSelf      = 1 // 京东配送
	ShipmentTypeToThird   = 2 // 京配转三方配送
	ShipmentTypeThird     = 3 // 第三方配送
	ShipmentTypeNormal    = 4 // 普通快递配送
	ShipmentTypeUnsupport = 9 // 不支持配送
)

const (
	OrderStatusSubmitSuccess = 0    // 提单成功
	OrderStatusWaitPay       = 1    // 等待付款
	OrderStatusPayed         = 4    // 已支付
	OrderStatusPrinting      = 6    // 等待打印，此过程执行预分拣、发票等逻辑。
	OrderStatusPicked        = 7    // 拣货完成
	OrderStatusOutStorage    = 8    // 出库完成
	OrderStatusWaitReceive   = 15   // 待用户确认收货，商品已经上车，送完站点
	OrderStatusRejectReceive = 16   // 用户拒收
	OrderStatusSigned        = 18   // 用户签收，订单正向流程生命周期完成
	OrderStatusLocked        = 21   // 订单锁定，系统中间状态
	OrderStatusCancelled     = -100 // 已取消
)

const (
	OrderTypeJdSelf   = 0  // 京东自营订单或混单
	OrderTypeShopSend = 1  // 商家自发货订单
	OrderTypeFactory  = 2  // 厂直订单
	OrderTypeOther    = 99 // 其他订单
)

const (
	PaymentOnline = 2 // 在线支付
)

type JdCommonResponse struct {
	ErrMsg     string `json:"errMsg"`
	ErrCode    int    `json:"errCode"`
	Success    bool   `json:"success"`
	SubErrCode string `json:"subErrCode,omitempty"`
	SubErrMsg  string `json:"subErrMsg,omitempty"`
}

type JdCommonErrResponse struct {
	Code          string `json:"code"`
	ErrorMessage  string `json:"errorMessage"`
	ErrorSolution string `json:"errorSolution"`
}

type SkuQuantityItem struct {
	Quantity int     `json:"quantity"`
	SkuId    string  `json:"skuId"`
	SkuPrice float64 `json:"skuPrice,omitempty"`
	SkuName  string  `json:"skuName,omitempty"`
}

type DeliveryAddress struct {
	ProvinceId  int64  `json:"provinceId"`
	CityId      int64  `json:"cityId"`
	CountyId    int64  `json:"countyId"`
	TownId      int64  `json:"townId"`
	FullAddress string `json:"fullAddress"`
}

type FreightFeeReq struct {
	PaymentType int                `json:"paymentType"`
	SkuList     []*SkuQuantityItem `json:"skuList"`
	Pin         string             `json:"pin"`
	OrderFee    float64            `json:"orderFee"`
	Address     DeliveryAddress    `json:"address"`
}

type FreightFeeResp struct {
	FreightFee float64 `json:"freightFee"`
}

type ShipmentTypeReq struct {
	PaymentType int                `json:"paymentType" desc:"支付方式。固定传2：在线支付"`
	SkuList     []*SkuQuantityItem `json:"skuList"`
	Pin         string             `json:"pin"`
	Address     DeliveryAddress    `json:"address"`
}

type ShipmentTypeResp struct {
	ShipmentType     int                 `json:"shipmentType"`
	ShipmentInfoList []*ShipmentInfoItem `json:"shipmentInfoList"`
}

type ShipmentInfoItem struct {
	SkuId          string         `json:"skuId"`
	ShipmentDetail ShipmentDetail `json:"shipmentDetail"`
}

type ShipmentDetail struct {
	ShipmentType   int            `json:"shipmentType"`
	AttachmentList []*ShipmentSku `json:"attachmentList"`
	GiftList       []*ShipmentSku `json:"giftList"`
}

type ShipmentSku struct {
	SkuId        string `json:"skuId"`
	ShipmentType int    `json:"shipmentType"`
}

type SubmitOrderReq struct {
	ChannelOrderId     string          `json:"channelOrderId"`
	Pin                string          `json:"pin"`
	ProductList        []*ProductItem  `json:"productList"`
	OrderFee           float64         `json:"orderFee" desc:"商品总金额，单位：元（不包含运费）"`
	FreightFee         float64         `json:"freightFee" desc:"运费金额 单位：元"`
	Address            DeliveryAddress `json:"address" desc:"收货地址信息"`
	Receiver           Receiver        `json:"receiver" desc:"收货人信息"`
	PaymentType        int             `json:"paymentType" desc:"支付方式。固定传2：在线支付"`
	ShipmentType       int             `json:"shipmentType" desc:"配送方式。1：京东配送 2：京配转三方配送 3：第三方配送 4：普通快递配送"`
	ChannelOrderSource string          `json:"channelOrderSource" desc:"订单来源。 JD：京东，DD：当当，PP：拍拍，QQ：QQ网购，SN：苏宁，GM：国美，WPH：唯品会，1688：阿里巴巴，POS：POS门店，TB：淘宝，TM：天猫，KS：快手，DY：抖音，OTHER：其他"`
	SendGoods          int             `json:"sendGoods" desc:"物品类型。固定传1"`
	Invoice            *Invoice        `json:"invoice,omitempty"`
	UserIp             string          `json:"userIp"`
	AutoCancelTime     int             `json:"autoCancelTime,omitempty"`
	DiscountFee        float64         `json:"discountFee,omitempty"`
}

type Receiver struct {
	ReceiverName   string `json:"receiverName"`
	ReceiverMobile string `json:"receiverMobile"`
	ReceiverEmail  string `json:"receiverEmail,omitempty"`
	ZipCode        string `json:"zipCode,omitempty"`
}

type ProductItem struct {
	MainSku SkuQuantityItem `json:"mainSku"`
}

type Invoice struct {
	InvoiceType       int                `json:"invoiceType" desc:"发票类型。2：增票，3：电子票 （如果选择开票此项必传）"`
	VatInvoice        *VatInvoice        `json:"vatInvoice,omitempty"`
	ElectronicInvoice *ElectronicInvoice `json:"electronicInvoice,omitempty"`
}

type VatInvoice struct {
	CompanyName     string      `json:"companyName" desc:"增票资质公司名称"`
	Code            string      `json:"code" desc:"增票资质公司纳税人标识号"`
	RegAddr         string      `json:"regAddr" desc:"增票资质公司注册地址"`
	RegPhone        string      `json:"regPhone" desc:"增票资质公司注册电话"`
	RegBank         string      `json:"regBank" desc:"增票资质公司注册银行"`
	RegBankAccount  string      `json:"regBankAccount" desc:"增票资质公司注册银行帐号"`
	ConsigneeName   string      `json:"consigneeName" desc:"增票邮寄收件人姓名"`
	ConsigneeMobile string      `json:"consigneeMobile" desc:"增票邮寄收件人电话"`
	VatAddress      *VatAddress `json:"vatAddress"`
}

type VatAddress struct {
	VatProvinceId  int    `json:"vatProvinceId"`
	VatCityId      int    `json:"vatCityId"`
	VatCountyId    int    `json:"vatCountyId"`
	VatTownId      int    `json:"vatTownId"`
	VatFullAddress string `json:"vatFullAddress"`
}

type ElectronicInvoice struct {
	SelectedInvoiceTitle  int    `json:"selectedInvoiceTitle" desc:"发票抬头 。 4：个人 5：公司"`
	ElectCompanyName      string `json:"electCompanyName" desc:"公司名称（如果发票抬头为：公司则此项必传）"`
	ElectCode             string `json:"electCode" desc:"纳税人识别号（如果发票抬头为：公司则此项必传）"`
	InvoiceConsigneeEmail string `json:"invoiceConsigneeEmail" desc:"联系邮箱（可用于发票信息的接收）"`
	InvoiceConsigneePhone string `json:"invoiceConsigneePhone" desc:"联系电话"`
}

type SubmitOrderResp struct {
	ChannelOrderId string `json:"channelOrderId"`
	OrderId        int64  `json:"orderId"`
	ChannelId      int64  `json:"channelId"`
	ExtendInfo     struct {
		Code string `json:"code"`
	} `json:"extendInfo"`
}

type QuerySubmitOrderReq struct {
	*Protocol
	ChannelOrderId string `json:"channelOrderId"`
	Pin            string `json:"pin"`
}

type QueryOrderDetailReq struct {
	*Protocol
	Pin     string `json:"pin"`
	OrderId int64  `json:"orderId"`
}

type QueryOrderDetailResp struct {
	OrderID          string           `json:"orderId"`
	BaseOrderInfo    BaseOrderInfo    `json:"baseOrderInfo"`
	OrderRelationFee OrderRelationFee `json:"orderRelationFee"`
	Shipment         Shipment         `json:"shipment"`
	SkuList          []*OrderSkuItem  `json:"skuList"`
}

type OrderSkuItem struct {
	ImgURL      string `json:"imgUrl"`
	Weight      string `json:"weight"`
	Color       string `json:"color"`
	CategoryID  string `json:"categoryId"`
	SkuID       string `json:"skuId"`
	SkuName     string `json:"skuName"`
	ShouldPrice string `json:"shouldPrice" desc:"商品单价，单位：元"`
	Quantity    string `json:"quantity"`
	Bulk        string `json:"bulk" desc:"商品体积"`
	SkuGiftType string `json:"skuGiftType" desc:"sku类型 1：主品；2：赠品"`
	MainSkuID   string `json:"mainSkuId"`
}

type BaseOrderInfo struct {
	RootOrderID string `json:"rootOrderId"`
	OrderStatus string `json:"orderStatus" desc:"订单状态 -100：已取消，0：提单成功，1：等待付款，4：已支付，6：等待打印，7：拣货完成，8：出库完成，15：待用户确认收货，16：用户拒收，18：用户签收，21：订单锁定，9：等待发货；"`

	SubmitTime       string `json:"submitTime"`
	CompleteTime     string `json:"completeTime"`
	PayTime          string `json:"payTime"`
	OutWarehouseTime string `json:"outWarehouseTime"`
	PaymentType      string `json:"paymentType"`
	Remark           string `json:"remark"`
	OrderType        string `json:"orderType" desc:"订单类型。0：京东自营订单或混单；1：商家自发货订单(POP-SOP)；2：厂直订单；99：其他订单"`
	VirtualSkuID     string `json:"virtualSkuId"`
}

type OrderRelationFee struct {
	ShouldPaymentFee string `json:"shouldPaymentFee" desc:"订单总金额，单位：元（包含运费）"`
	FreightFee       string `json:"freightFee"`
	DiscountFee      string `json:"discountFee"`
}

type Shipment struct {
	ShipmentType string `json:"shipmentType" desc:"配送方式。70：快递运输"`
}

type ConfirmDeliveryReq struct {
	*Protocol
	ClientPort string `json:"clientPort"`
	OrderId    int64  `json:"orderId"`
}

type PushOrderReq struct {
	*Protocol
	OrderId int64 `json:"orderId"`
}

type OrderPayInfoResp struct {
	RefundDetailList []*RefundDetail `json:"refundDetailList"`
	RefundTotalFee   float64         `json:"refundTotalFee"`
	RootOrderId      int64           `json:"rootOrderId"`
	PaidInDetailList []*PaidInDetail `json:"paidInDetailList"`
	ChannelOrderId   string          `json:"channelOrderId"`
	FreightFee       float64         `json:"freightFee"`
	PaidInTotalFee   float64         `json:"paidInTotalFee"`
	OrderFee         float64         `json:"orderFee"`
}

type PaidInDetail struct {
	PaidInFee  float64 `json:"paidInFee"`
	PaidInType int     `json:"paidInType"`
	PaidInTime string  `json:"paidInTime"`
}

type RefundDetail struct {
	RefundFee  float64 `json:"refundFee"`
	RefundType int     `json:"refundType"`
	RefundTime string  `json:"refundTime"`
}

type InvoiceDetail struct {
	OrderId              int64         `json:"orderId"`
	TotalPrice           float64       `json:"totalPrice"`
	IvcTime              string        `json:"ivcTime"`
	IvcNo                string        `json:"ivcNo"`
	InvoiceSkuDetailList []*InvoiceSku `json:"invoiceSkuDetailList"`
	UpdateTime           string        `json:"updateTime"`
	Remark               string        `json:"remark"`
	IvcContentName       string        `json:"ivcContentName"`
	Valid                string        `json:"valid"`
	TaxRate              float64       `json:"taxRate"`
	IvcType              int           `json:"ivcType" desc:"发票类型 1：普票；2：增票；3：电子发票"`
	TotalTaxPrice        float64       `json:"totalTaxPrice"`
	IvcTitle             string        `json:"ivcTitle"`
	IvcContentType       int           `json:"ivcContentType" desc:"发票内容类型。1：明细；2：类别"`
	FileUrl              string        `json:"fileUrl"`
	ExpandColumn         string        `json:"expandColumn"`
	IvcCode              string        `json:"ivcCode"`
}

type InvoiceSku struct {
	SkuName    string  `json:"skuName"`
	TaxRate    float64 `json:"taxRate"`
	IvcSkuName string  `json:"ivcSkuName"`
	Unit       int     `json:"unit"`
	IvcPrice   float64 `json:"ivcPrice"`
	Isn        string  `json:"isn"`
	Price      float64 `json:"price"`
	Num        int     `json:"num"`
	IvcNum     int     `json:"ivcNum"`
	SkuId      int64   `json:"skuId"`
}

type CancelOrderParams struct {
	*Protocol
	Pin string `json:"pin"`
	CancelOrderReq
}

type CancelOrderReq struct {
	OrderId          int64 `json:"orderId"`
	CancelReasonCode int   `json:"cancelReasonCode" validate:"oneof=1 2 3 4 5 6 7 100" oneof:"请选择售后原因" desc:"取消原因。1：订单不能按预计时间送达 2：其他渠道价格更低 3：该商品降价了 4：不想买了 5：其他原因 6：操作有误（商品、地址等选错） 7：商品无货 100：其他"`
	CancelReasonType int   `json:"cancelReasonType" desc:"取消类型。1：未支付取消 2：用户取消 3：风控取消"`
	CancelType       int   `json:"cancelType" desc:"取消类型。1：订单取消；2：订单拦截"`
}

type CancelOrderResp struct {
	CancelStatus int `json:"cancelStatus" desc:"订单取消状态。同步接入取消申请，则状态1 ：取消失败 3：取消成功；异步接入取消申请，则状态 1：取消申请失败 3：取消申请成功（通过订单取消MQ获取取消是否成功的结果）"`
}

type StateSkuQuantityItem struct {
	Quantity int     `json:"quantity"`
	SkuId    int64   `json:"skuId"`
	SkuPrice float64 `json:"skuPrice,omitempty"`
	SkuName  string  `json:"skuName,omitempty"`
}

type StockStateReq struct {
	SkuQuantityList []*StateSkuQuantityItem `json:"skuQuantityList"`
	Address         DeliveryAddress         `json:"address"`
}

type StockStateRespItem struct {
	AreaStockState int                  `json:"areaStockState" desc:"0-无货 1-有货（现货） 2-采购中"`
	LeadTime       string               `json:"leadTime" desc:"采购中商品的预计到货时间"`
	SkuQuantity    StateSkuQuantityItem `json:"skuQuantity"`
}

type AddressData struct {
	ID       int64  `json:"id" desc:"子地址id"`
	ParentId int64  `json:"parentId" desc:"父地址id（同入参）"`
	Name     string `json:"name" desc:"子地址名称"`
}

type LogisticTrace struct {
	WaybillCode      string               `json:"waybillCode"`
	LogisticsName    string               `json:"logisticsName"`
	OperatorNodeList []*LogisticTraceNode `json:"operatorNodeList"`
}

// LogisticTraceNode 轨迹节点
type LogisticTraceNode struct {
	Content        string `json:"content"`
	GroupState     string `json:"groupState"`
	ScanState      string `json:"scanState"`
	MsgTime        int64  `json:"msgTime"`
	SystemOperator string `json:"systemOperator"`
	OrderId        int64  `json:"orderId"`
}

type QueryOrderLogisticsReq struct {
	Protocol
	Pin     string `json:"pin"`
	OrderId int64  `json:"orderId"`
}

type QuerySkuListReq struct {
	PageSize int64  `json:"pageSize"`
	ScrollId string `json:"scrollId"`
}

type QuerySkuListResp struct {
	ScrollId string                   `json:"scrollId"`
	Total    int64                    `json:"total"`
	Entries  []*QuerySkuListRespEntry `json:"entries"`
}

type QuerySkuListRespEntry struct {
	SkuId         int64  `json:"skuId"`
	SkuName       string `json:"skuName"`
	OuterId       string `json:"outerId"`
	ImgUrl        string `json:"imgUrl"`
	CategoryId1   int64  `json:"categoryId1"`
	CategoryName1 string `json:"categoryName1"`
	CategoryId2   int64  `json:"categoryId2"`
	CategoryName2 string `json:"categoryName2"`
	CategoryId    int64  `json:"categoryId"`
	CategoryName  string `json:"categoryName"`
	BrandName     string `json:"brandName"`
	BrandId       int64  `json:"brandId"`
	SkuStatus     int64  `json:"skuStatus"`
	Modified      int64  `json:"modified"`
	Created       int64  `json:"created"`
	EnBrandName   string `json:"enBrandName"`
	BrandCountry  string `json:"brandCountry"`
	WareType      int    `json:"wareType"`
}

type QueryBrotherSkuListReq struct {
	SkuIds []int64 `json:"skuIdSet"`
}

type QueryBrotherSkuListItem struct {
	IsSuccess     bool    `json:"isSuccess"`
	ErrorMessage  string  `json:"errorMessage"`
	SkuId         int64   `json:"skuId"`
	BrotherSkuIds []int64 `json:"brotherSkuIds"`
}

type QuerySkuDetailReq struct {
	DetailAssemblyType int64   `json:"detailAssemblyType"`
	SkuIdSet           []int64 `json:"skuIdSet"`
}

type QuerySkuDetailResp struct {
	ImageInfos      []ImageInfo     `json:"imageInfos"`
	WReadMe         string          `json:"wReadMe"`
	Specifications  []Specification `json:"specifications"`
	SkuID           int64           `json:"skuId"`
	ExtAtts         []ExtAtt        `json:"extAtts"`
	SkuBaseInfo     SkuBaseInfo     `json:"skuBaseInfo"`
	SkuBigFieldInfo SkuBigFieldInfo `json:"skuBigFieldInfo" desc:"大字段信息"`
}

type ImageInfo struct {
	Path      string `json:"path"`
	Features  string `json:"features"`
	OrderSort int64  `json:"orderSort"`
	IsPrimary int64  `json:"isPrimary"`
	Position  int64  `json:"position"`
	Type      int64  `json:"type"`
}

type Specification struct {
	GroupName  string      `json:"groupName"`
	Attributes []Attribute `json:"attributes"`
}

type Attribute struct {
	AttName  string   `json:"attName"`
	ValNames []string `json:"valNames"`
}

type ExtAtt struct {
	AttName  string   `json:"attName"`
	ValNames []string `json:"valNames"`
}

type SkuBaseInfo struct {
	SkuName       string `json:"skuName" desc:""`
	VenderName    string `json:"venderName" desc:""`
	ShopName      string `json:"shopName" desc:""`
	CategoryId1   int64  `json:"categoryId1" desc:""`
	CategoryId2   int64  `json:"categoryId2" desc:""`
	CategoryId    int64  `json:"categoryId" desc:""`
	Unit          string `json:"unit" desc:"销售单位"`
	ProductId     int64  `json:"productId" desc:"商品编号"`
	SkuStatus     int64  `json:"skuStatus" desc:"	上下架状态"`
	Yn            int64  `json:"yn" desc:"是否有效。0：无效 1：有效"`
	Fare          int64  `json:"fare" desc:"运费模板id"`
	CategoryName1 string `json:"categoryName1" desc:"一级分类名称"`
	CategoryName2 string `json:"categoryName2" desc:"二级分类名称"`
	CategoryName  string `json:"categoryName" desc:"三级分类名称"`
	SkuInfoType   int64  `json:"skuInfoType" desc:"sku信息类型 1-图书 2-音像"`
	WareType      int64  `json:"wareType" desc:"商品类型 （1-普通商品; 2-虚拟组套）"`
}

type SkuBigFieldInfo struct {
	PcWdis        string `json:"pcWdis" desc:"PC端商品介绍"`
	PcHtmlContent string `json:"pcHtmlContent" desc:"PC HTML标签内容"`
	PcCssContent  string `json:"pcCssContent" desc:"PC css样式内容"`
	PcJsContent   string `json:"pcJsContent" desc:"PC js内容"`
}

type QuerySkuPriceReq struct {
	SkuIdSet []int64 `json:"skuIdSet"`
}

type QuerySkuPriceResp struct {
	CustomerId   int64      `json:"customerId" desc:"客户id，客户唯一身份标识"`
	ChannelId    int64      `json:"channelId" desc:"渠道id，客户对应的细分渠道标识"`
	SkuPriceList []SkuPrice `json:"skuPriceList"`
}

type SkuPrice struct {
	IsSuccess       bool    `json:"isSuccess" desc:"是否获取价格成功"`
	ChannelId       string  `json:"errorMessage" desc:"若sku价格查询失败（isSuccess!=true），则返回原因"`
	SkuPrice        float64 `json:"skuPrice" desc:"sku价格，仅获取价格成功时(isSuccess=true)有值,单位:元 小数点后两位"`
	SkuId           int64   `json:"skuId" desc:"京东商品skuId"`
	PriceTypeMark   string  `json:"priceTypeMark" desc:"商品价格类型标识(1：基础价格，2：促销价格)"`
	ProfitRate      float64 `json:"profitRate" desc:"商品利润率"`
	BackStagePrice  string  `json:"backStagePrice" desc:"京东后台价"`
	PriceUpdateTime int64   `json:"priceUpdateTime" desc:"商品价格更新时间"`
}

type CanApplyInfoReq struct {
	Pin           string `json:"pin"`
	SkuId         int64  `json:"skuId"`
	OrderId       int64  `json:"orderId"`
	AfsDetailType int    `json:"afsDetailType" desc:"商品类型：10-主品，20-赠品"`
}

type CanApplyInfoItem struct {
	AfsSupportedTypes []*AfsSupportedTypeItem `json:"afsSupportedTypes"`
	CanApply          int                     `json:"canApply" desc:"该订单是否可申请售后 0：不可申请 1：可申请"`
	OrderId           int64                   `json:"orderId" desc:"京东订单号"`
	AppliedNum        int                     `json:"appliedNum" desc:"已经申请售后的商品数量"`
	CannotApplyTip    string                  `json:"cannotApplyTip" desc:"不可申请提示"`
	SkuId             int64                   `json:"skuId" desc:"京东商品编号"`
}

type AfsSupportedTypeItem struct {
	AfsTypeName string `json:"afsTypeName" desc:"售后服务类型名称"`
	AfsType     int    `json:"afsType" desc:"售后服务类型编号"`
}

type QueryApplyReasonsReq struct {
	CanApplyInfoReq
	OrderId int64 `json:"orderId" desc:"售后服务单对应的京东订单号"`
}

type ApplyReasonItem struct {
	ApplyReasonName string `json:"applyReasonName" desc:"售后申请问题描述文字"`
	ApplyReasonId   int    `json:"applyReasonId" desc:"售后申请原因ID"`
	NeedUploadPic   bool   `json:"needUploadPic" desc:"是否必须上传图片"`
}

type CreateAfsApplyReq struct {
	ApplyReasonName   string          `json:"applyReasonName"`
	ApplyReasonId     int             `json:"applyReasonId" desc:"售后申请原因ID"`
	Pin               string          `json:"pin"`
	ChannelAfsApplyId string          `json:"channelAfsApplyId" desc:"渠道售后服务单申请单号"`
	AfsType           int             `json:"afsType" desc:"用户期望的售后服务类型 10：退货 20：换货"`
	QuestionPic       string          `json:"questionPic,omitempty" desc:"售后申请问题描述图片，多个时用逗号分隔"`
	OrderId           int64           `json:"orderId"`
	SkuQuantity       SkuQuantity     `json:"skuQuantity"`
	PickWareType      int             `json:"pickWareType" desc:"取件方式上门取件NORMAL_PICKWARE(4)大家电上门取件MAJOR_NORMAL_PICKWARE(8)客户发货CUSTOMER_SEND_WARE(40)"`
	PickWareAddress   PickWareAddress `json:"pickWareAddress"`
}

type SkuQuantity struct {
	SkuId         int64  `json:"skuId"`
	SkuName       string `json:"skuName"`
	Quantity      int    `json:"quantity"`
	AfsDetailType int    `json:"afsDetailType"`
}

type PickWareAddress struct {
	ProvinceId    int    `json:"provinceId"`
	CityId        int    `json:"cityId"`
	CountyId      int    `json:"countyId"`
	TownId        int    `json:"townId"`
	FullAddress   string `json:"fullAddress"`
	AddressDetail string `json:"addressDetail"`
}

type CreateAfsApplyResp struct {
	ChannelAfsApplyId string `json:"channelAfsApplyId" desc:"渠道售后服务单申请单号"`
	AfsApplyId        int64  `json:"afsApplyId" desc:"京东售后服务单申请单号"`
}

type GetAfsLogisticAddressResp struct {
	Address         string `json:"address"`
	ContactsZipCode string `json:"contactsZipCode"`
	ContactsName    string `json:"contactsName"`
	ContactsMobile  string `json:"contactsMobile"`
}

type PostBackLogisitcBillReq struct {
	Pin              string `json:"pin"`
	AfsServiceId     int64  `json:"afsServiceId"`
	LogisticsCompany string `json:"logisticsCompany"`
	WaybillCode      string `json:"waybillCode"`
	SendGoodsDate    string `json:"sendGoodsDate"`
}

type PostBackLogisitcBillResp struct {
	PostBackResult bool   `json:"postBackResult"`
	Message        string `json:"message"`
}

type AfsServiceDetailResp struct {
	AfsTypeName         string                       `json:"afsTypeName"`
	OrderId             int64                        `json:"orderId"`
	ApplyReasonId       int                          `json:"applyReasonId"`
	ProcessNotes        string                       `json:"processNotes" desc:"售后服务单处理意见"`
	AfsType             int                          `json:"afsType"`
	AfsServiceStep      int                          `json:"afsServiceStep"`
	DesenCustomerMobile string                       `json:"desen_customerMobile"`
	ProcessedDate       int64                        `json:"processedDate"`
	Pin                 string                       `json:"pin"`
	ApproveResult       int                          `json:"approveResult"`
	ApplyReasonName     string                       `json:"applyReasonName"`
	CustomerEmail       string                       `json:"customerEmail"`
	AfsServiceState     int                          `json:"afsServiceState" desc:"服务单状态"`
	ApproveResultName   string                       `json:"approveResultName" desc:"售后服务单审核结果"`
	ApproveNotes        string                       `json:"approveNotes" desc:"售后服务单审核意见"`
	ProcessResult       int                          `json:"processResult" desc:"售后服务单审核时间"`
	AfsServiceId        int64                        `json:"afsServiceId"`
	AfsApplyTime        int64                        `json:"afsApplyTime"`
	AfsApplyId          int64                        `json:"afsApplyId"`
	CustomerName        string                       `json:"customerName"`
	ApprovedDate        int64                        `json:"approvedDate"`
	AfsServiceStateName string                       `json:"afsServiceStateName"`
	ProcessResultName   string                       `json:"processResultName" desc:"售后服务单处理结果"`
	AfsServiceStepName  string                       `json:"afsServiceStepName" desc:"售后服务单状态"`
	NewOrderId          int64                        `json:"newOrderId" desc:"售后换新单订单号"`
	CustomerMobile      string                       `json:"customerMobile" desc:"用户联系方式（售后联系人联系方式）"`
	QuestionPic         string                       `json:"questionPic" desc:"售后申请问题描述图片"`
	ReturnWareType      int                          `json:"returnWareType" desc:"售后返件类型 1：客户发货"`
	SkuQuantity         *AfsServiceDetailSkuQuantity `json:"skuQuantity"`
}

type AfsServiceDetailSkuQuantity struct {
	SkuName      string `json:"skuName"`
	Quantity     int    `json:"quantity"`     //退换货数量
	SkuType      int    `json:"skuType"`      //标识商品属性1单品、2买赠：赠品套装中的主商品、3买赠：赠品套装中的赠品
	ValidNumFlag int    `json:"validNumFlag"` //赠品申请标识 1代表申请了，0代表没申请或释放了，后续可以继续申请
	SkuId        int64  `json:"skuId"`
}
