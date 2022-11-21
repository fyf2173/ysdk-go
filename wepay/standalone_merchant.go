package wepay

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
)

type StandaloneMerchantConfig struct {
	AppId   string
	AppName string
	MchId   string

	SerialNum   string
	PrivateKey  *rsa.PrivateKey
	Certificate *x509.Certificate
	ApiV3Key    string
}

type StandaloneMerchantClient struct {
	StandaloneMerchantConfig
	cli *core.Client
}

func (mc *StandaloneMerchantClient) ClientInstance() *core.Client {
	return mc.cli
}

// NewStandaloneMerchantClient 初始化独立商户客户端
func NewStandaloneMerchantClient(ctx context.Context, conf StandaloneMerchantConfig) *StandaloneMerchantClient {
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(conf.MchId, conf.SerialNum, conf.PrivateKey, conf.ApiV3Key),
	}
	cli, err := core.NewClient(ctx, opts...)
	if err != nil {
		panic(err)
	}
	return &StandaloneMerchantClient{StandaloneMerchantConfig: conf, cli: cli}
}

func (mc *StandaloneMerchantClient) NewPrepayRequest(req *CreateOrderReq) *jsapi.PrepayRequest {
	return &jsapi.PrepayRequest{
		Appid:       core.String(mc.AppId),
		Mchid:       core.String(mc.MchId),
		Description: core.String(req.Description),
		OutTradeNo:  core.String(req.OutTradeNo),
		TimeExpire:  core.Time(req.TimeExpire),
		NotifyUrl:   core.String(req.NotifyUrl),
		GoodsTag:    core.String(req.GoodsTag),
		LimitPay:    req.LimitPay,
		Amount: &jsapi.Amount{
			Total:    core.Int64(req.TotalAmount),
			Currency: core.String("CNY"),
		},
		Payer: &jsapi.Payer{
			Openid: core.String(req.PayerOpenId),
		},
		Detail: &jsapi.Detail{
			CostPrice: core.Int64(req.OriginAmount),
		},
	}
}

// Prepay 预支付
func (mc *StandaloneMerchantClient) Prepay(ctx context.Context, req *jsapi.PrepayRequest) (resp *jsapi.PrepayResponse, result *core.APIResult, err error) {
	svc := jsapi.JsapiApiService{Client: mc.ClientInstance()}
	return svc.Prepay(ctx, *req)
}
