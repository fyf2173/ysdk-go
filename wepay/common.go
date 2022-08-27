package wepay

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"strconv"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/partnerpayments/jsapi"
	jspay "github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

type MerchantConfig struct {
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

type MerchantClient struct {
	MerchantConfig
	cli *core.Client
}

func (mc *MerchantClient) Get() *core.Client {
	return mc.cli
}

// NewMerchantClient 初始化商户客户端
func NewMerchantClient(ctx context.Context, conf MerchantConfig) *MerchantClient {
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(conf.SpMchId, conf.SpSerialNum, conf.SpPrivateKey, conf.SpApiV3Key),
	}
	cli, err := core.NewClient(ctx, opts...)
	if err != nil {
		panic(err)
	}
	return &MerchantClient{MerchantConfig: conf, cli: cli}
}

// NewPaymentResp 生成支付数据包
func (mc *MerchantClient) NewPaymentResp(ctx context.Context, prepayId string) (*jspay.PrepayWithRequestPaymentResponse, error) {
	paymentResp := new(jspay.PrepayWithRequestPaymentResponse)
	paymentResp.PrepayId = core.String(prepayId)
	paymentResp.SignType = core.String("RSA")
	paymentResp.Appid = core.String(mc.SubAppId)
	paymentResp.TimeStamp = core.String(strconv.FormatInt(time.Now().Unix(), 10))
	nonce, err := utils.GenerateNonce()
	if err != nil {
		return nil, fmt.Errorf("prepay generate nonce error %+v", err)
	}
	paymentResp.NonceStr = core.String(nonce)
	paymentResp.Package = core.String("prepay_id=" + prepayId)
	message := fmt.Sprintf("%s\n%s\n%s\n%s\n", *paymentResp.Appid, *paymentResp.TimeStamp, *paymentResp.NonceStr, *paymentResp.Package)
	signatureResult, err := mc.Get().Sign(ctx, message)
	if err != nil {
		return nil, fmt.Errorf("prepay sign error %+v", err)
	}
	paymentResp.PaySign = core.String(signatureResult.Signature)
	return paymentResp, nil
}

// Prepay 预支付
func (mc *MerchantClient) Prepay(ctx context.Context, req *jsapi.PrepayRequest) (resp *jsapi.PrepayResponse, result *core.APIResult, err error) {
	svc := jsapi.JsapiApiService{Client: mc.Get()}
	return svc.Prepay(ctx, *req)
}

type RefundReq struct {
	TransactionId string
	OutTradeNo    string
	OutRefundNo   string
	Reason        string
	NotifyUrl     string
	Refund        int64
	Total         int64
}

func (mc *MerchantClient) NewRefundsApply(req RefundReq) *refunddomestic.CreateRequest {
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
func (mc *MerchantClient) ApplyRefunds(ctx context.Context, req *refunddomestic.CreateRequest) (resp *refunddomestic.Refund, result *core.APIResult, err error) {
	svc := refunddomestic.RefundsApiService{Client: mc.Get()}
	return svc.Create(ctx, *req)
}

const (
	NotifyPayed  = "payed"
	NotifyRefund = "refund"

	TransactionSuccess = "TRANSACTION.SUCCESS"
	RefundSuccess      = "REFUND.SUCCESS"
)

// NotifyHandler 通知回调处理
func (mc *MerchantClient) NotifyHandler(ctx context.Context) (*notify.Handler, error) {
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
