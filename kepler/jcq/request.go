package jcq

import (
	"encoding/json"
	"fmt"

	sdkHttp "github.com/fyf2173/ysdk-go/http"
)

// Request 消费/确认
func (jc *Client) Request(method string, path string, params interface{}, out interface{}) error {
	jh := jc.NewHeader()
	jh.Signature = jc.GetSignature(jh.GetSignSource(params))

	opts := []sdkHttp.Option{
		sdkHttp.SetRequestHeader("accessKey", jh.AccessKey),
		sdkHttp.SetRequestHeader("dateTime", jh.DateTime),
		sdkHttp.SetRequestHeader("signature", jh.Signature),
	}
	var resp CommonConsumerResp
	if err := sdkHttp.Request(method, fmt.Sprintf("%s%s", endPoint, path), params, &resp, opts...); err != nil {
		return err
	}
	if resp.Error != nil && resp.Error.Code != 0 {
		return fmt.Errorf("code=%d,message=%s,status=%s", resp.Error.Code, resp.Error.Message, resp.Error.Status)
	}
	if out == nil {
		return nil
	}
	return json.Unmarshal(resp.Result, &out)
}
