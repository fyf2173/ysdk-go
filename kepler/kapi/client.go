package kapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	sdkHttp "github.com/fyf2173/ysdk-go/http"

	"github.com/fyf2173/ysdk-go/crypto"
)

type Client struct {
	Pin        string
	ChannelId  int64
	CustomerID int64
	AppID      string
	AppKey     string
	AppSecret  string
	Token      string
}

func (kc *Client) GetProtocolParams() *Protocol {
	return &Protocol{
		CustomerId: kc.CustomerID,
		ChannelId:  kc.ChannelId,
		AppKey:     kc.AppKey,
		ClientIp:   "127.0.0.1",
		OpName:     kc.Pin,
	}
}

func (kc *Client) Sign(method string, params interface{}) string {
	var (
		signMap   = make(map[string]string)
		paramsMap = make(map[string]interface{})
	)

	signMap["method"] = method
	signMap["access_token"] = kc.Token
	signMap["app_key"] = kc.AppKey
	signMap["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	signMap["format"] = "json"
	signMap["v"] = "2.0"

	pb, _ := json.Marshal(params)
	_ = json.Unmarshal(pb, &paramsMap)
	b, _ := json.Marshal(paramsMap)

	signMap["360buy_param_json"] = string(b)

	signKeys := make([]string, 0)
	for sk := range signMap {
		signKeys = append(signKeys, sk)
	}
	sort.Strings(signKeys)
	var buffer bytes.Buffer
	for _, signKey := range signKeys {
		if signv := signMap[signKey]; len(signv) > 0 {
			buffer.Write([]byte(signKey))
			buffer.Write([]byte(signv))
		}
	}
	signStr := kc.AppSecret + buffer.String() + kc.AppSecret
	signMap["sign"] = strings.ToUpper(crypto.Md5Str(signStr))
	vals := url.Values{}
	for k, v := range signMap {
		if v != "" {
			vals.Add(k, v)
		}
	}
	return vals.Encode()
}

func (kc *Client) transRespName(method string) string {
	var respNameArr = strings.Split(method, ".")
	respNameArr = append(respNameArr, "responce")
	return strings.Join(respNameArr, "_")
}

type BaseResponse struct {
	Result struct {
		JdCommonResponse
		Data           json.RawMessage `json:"data"`
		StockStateList json.RawMessage `json:"stockStateList"`
	} `json:"result"`
	Code   string `json:"code"`
	ZhDesc string `json:"zh_desc"`
	EnDesc string `json:"en_desc"`
}

func (kc *Client) Request(apiMethod string, params interface{}, response interface{}) error {
	var (
		resp          = make(map[string]BaseResponse)
		responseField = kc.transRespName(apiMethod)
		link          = RouterJson + "?" + kc.Sign(apiMethod, params)
	)

	if err := sdkHttp.Request(http.MethodGet, link, nil, &resp, sdkHttp.SetContentTypeForm); err != nil {
		return err
	}
	if errResp, ok := resp["error_response"]; ok {
		return fmt.Errorf("code=%s,msg=%s", errResp.Code, errResp.ZhDesc)
	}
	data, ok := resp[responseField]
	if !ok {
		return fmt.Errorf("JD返回数据异常")
	}

	if data.Result.Success == false || data.Result.ErrCode != http.StatusOK {
		return fmt.Errorf("%s", data.Result.ErrMsg)
	}
	if response == nil {
		return nil
	}
	if data.Result.StockStateList != nil {
		return json.Unmarshal(data.Result.StockStateList, &response)
	}
	return json.Unmarshal(data.Result.Data, &response)
}
