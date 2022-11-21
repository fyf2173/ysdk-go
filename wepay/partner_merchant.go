package wepay

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	partnerJsapi "github.com/wechatpay-apiv3/wechatpay-go/services/partnerpayments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"
)

type PartnerMerchantConfig struct {
	SpAppId    string
	SpMchId    string
	SubAppId   string
	SubAppName string
	SubMchId   string

	SpSerialNum   string
	SpPrivateKey  *rsa.PrivateKey
	SpCertificate *x509.Certificate
	SpApiV3Key    string
}

type PartnerMerchantClient struct {
	PartnerMerchantConfig
	cli *core.Client
}

func (mc *PartnerMerchantClient) ClientInstance() *core.Client {
	return mc.cli
}

// NewPartnerMerchantClient 初始化服务商户客户端
func NewPartnerMerchantClient(ctx context.Context, conf PartnerMerchantConfig) *PartnerMerchantClient {
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(conf.SpMchId, conf.SpSerialNum, conf.SpPrivateKey, conf.SpApiV3Key),
	}
	cli, err := core.NewClient(ctx, opts...)
	if err != nil {
		panic(err)
	}
	return &PartnerMerchantClient{PartnerMerchantConfig: conf, cli: cli}
}

func (mc *PartnerMerchantClient) NewPrepayRequest(req *CreateOrderReq) *partnerJsapi.PrepayRequest {
	return &partnerJsapi.PrepayRequest{
		SpAppid:     core.String(mc.SpAppId),
		SpMchid:     core.String(mc.SpMchId),
		SubAppid:    core.String(mc.SubAppId),
		SubMchid:    core.String(mc.SubMchId),
		Description: core.String(req.Description),
		OutTradeNo:  core.String(req.OutTradeNo),
		TimeExpire:  core.Time(req.TimeExpire),
		NotifyUrl:   core.String(req.NotifyUrl),
		GoodsTag:    core.String(req.GoodsTag),
		LimitPay:    req.LimitPay,
		Amount: &partnerJsapi.Amount{
			Total:    core.Int64(req.TotalAmount),
			Currency: core.String("CNY"),
		},
		Payer: &partnerJsapi.Payer{
			SubOpenid: core.String(req.PayerOpenId),
		},
		Detail: &partnerJsapi.Detail{
			CostPrice: core.Int64(req.OriginAmount),
		},
	}
}

// Prepay 预支付
func (mc *PartnerMerchantClient) Prepay(ctx context.Context, req *partnerJsapi.PrepayRequest) (resp *partnerJsapi.PrepayResponse, result *core.APIResult, err error) {
	svc := partnerJsapi.JsapiApiService{Client: mc.ClientInstance()}
	return svc.Prepay(ctx, *req)
}

func (mc *PartnerMerchantClient) NewRefundsApply(req RefundReq) *refunddomestic.CreateRequest {
	return &refunddomestic.CreateRequest{
		SubMchid:      core.String(mc.SubMchId),
		TransactionId: core.String(req.TransactionId),
		OutTradeNo:    core.String(req.OutTradeNo),
		OutRefundNo:   core.String(req.OutRefundNo),
		Reason:        core.String(req.Reason),
		NotifyUrl:     core.String(req.NotifyUrl),
		Amount: &refunddomestic.AmountReq{
			Currency: core.String("CNY"),
			Refund:   core.Int64(req.Refund),
			Total:    core.Int64(req.Total),
		},
	}
}

// ApplyRefunds 申请退款
func (mc *PartnerMerchantClient) ApplyRefunds(ctx context.Context, req *refunddomestic.CreateRequest) (resp *refunddomestic.Refund, result *core.APIResult, err error) {
	svc := refunddomestic.RefundsApiService{Client: mc.ClientInstance()}
	return svc.Create(ctx, *req)
}

const (
	NotifyPayed  = "payed"
	NotifyRefund = "refund"

	TransactionSuccess = "TRANSACTION.SUCCESS"
	RefundSuccess      = "REFUND.SUCCESS"
)

// NotifyHandler 通知回调处理
func (mc *PartnerMerchantClient) NotifyHandler(ctx context.Context) (*notify.Handler, error) {
	// 1. 使用 `RegisterDownloaderWithPrivateKey` 注册下载器
	err := downloader.MgrInstance().RegisterDownloaderWithPrivateKey(ctx, mc.SpPrivateKey, mc.SpSerialNum, mc.SpMchId, mc.SpApiV3Key)
	if err != nil {
		return nil, err
	}
	// 2. 获取商户号对应的微信支付平台证书访问器
	certificateVisitor := downloader.MgrInstance().GetCertificateVisitor(mc.SpMchId)
	// 3. 使用apiv3 key、证书访问器初始化 `notify.Handler`
	handler := notify.NewNotifyHandler(mc.SpApiV3Key, verifiers.NewSHA256WithRSAVerifier(certificateVisitor))
	return handler, nil
}
