package wepay

import (
	"context"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"
)

type RefundReq struct {
	TransactionId string
	OutTradeNo    string
	OutRefundNo   string
	Reason        string
	NotifyUrl     string
	Refund        int64
	Total         int64
}

func NewRefundsApply(mchId string, req RefundReq) *refunddomestic.CreateRequest {
	return &refunddomestic.CreateRequest{
		SubMchid:      core.String(mchId),
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
func ApplyRefunds(ctx context.Context, client *core.Client, req *refunddomestic.CreateRequest) (resp *refunddomestic.Refund, result *core.APIResult, err error) {
	svc := refunddomestic.RefundsApiService{Client: client}
	return svc.Create(ctx, *req)
}
