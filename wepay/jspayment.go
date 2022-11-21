package wepay

import (
	"context"
	"fmt"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"strconv"
	"time"
)

type CreateOrderReq struct {
	Description  string
	OutTradeNo   string
	Attach       string
	NotifyUrl    string
	GoodsTag     string
	TotalAmount  int64
	OriginAmount int64
	PayerOpenId  string
	LimitPay     []string
	TimeExpire   time.Time
}

// NewJsPayment 生成支付数据包
func NewJsPayment(ctx context.Context, client *core.Client, appId, prepayId string) (*jsapi.PrepayWithRequestPaymentResponse, error) {
	paymentResp := new(jsapi.PrepayWithRequestPaymentResponse)
	paymentResp.PrepayId = core.String(prepayId)
	paymentResp.SignType = core.String("RSA")
	paymentResp.Appid = core.String(appId)
	paymentResp.TimeStamp = core.String(strconv.FormatInt(time.Now().Unix(), 10))
	nonce, err := utils.GenerateNonce()
	if err != nil {
		return nil, fmt.Errorf("prepay generate nonce error %+v", err)
	}
	paymentResp.NonceStr = core.String(nonce)
	paymentResp.Package = core.String("prepay_id=" + prepayId)
	message := fmt.Sprintf("%s\n%s\n%s\n%s\n", *paymentResp.Appid, *paymentResp.TimeStamp, *paymentResp.NonceStr, *paymentResp.Package)
	signatureResult, err := client.Sign(ctx, message)
	if err != nil {
		return nil, fmt.Errorf("prepay sign error %+v", err)
	}
	paymentResp.PaySign = core.String(signatureResult.Signature)
	return paymentResp, nil
}
