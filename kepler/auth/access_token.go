package auth

import (
	"fmt"
	"net/http"

	sdkHttp "github.com/fyf2173/ysdk-go/http"
)

const (
	AccessToken  = "https://open-oauth.jd.com/oauth2/access_token?app_key=%s&app_secret=%s&grant_type=authorization_code&code=%s"
	RefreshToken = "https://open-oauth.jd.com/oauth2/refresh_token?app_key=%s&app_secret=%s&grant_type=refresh_token&refresh_token=%s"
)

type JdAccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	ExpiresAt    int64  `json:"expires_at"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	OpenID       string `json:"open_id"`
	Time         int64  `json:"time" desc:"13位时间戳"`

	Code      int    `json:"code,omitempty"`
	Msg       string `json:"msg,omitempty"`
	RequestId string `json:"requestId,omitempty"`
}

// AccessToken 获取token
// {"access_token":"9097002316914b3583f3499fa9ad7d69xota","expires_in":86400,"refresh_token":"ab87a46362ea4e7799b5ee1c4f7bba41n2e0","scope":"snsapi_base","open_id":"49L4y3KtuV87ymNjRTXCB0vnEiY45eyiYwHmCz4zSkU","uid":"0201092577","time":1659322812128,"token_type":"bearer","code":0,"xid":"o*AATzS6F5-Y2HM3BqNKz0ZQ9LMmZkNScOYrCptRUJNvvGC6j2970"}
func (ac *AccessClient) AccessToken(code string) (*JdAccessTokenResponse, error) {
	var response JdAccessTokenResponse
	link := fmt.Sprintf(AccessToken, ac.AppKey, ac.AppSecret, code)

	if err := sdkHttp.Request(http.MethodGet, link, nil, &response); err != nil {
		return nil, fmt.Errorf("query access token failed, err=%+v", err)
	}
	if response.Code != 0 {
		return nil, fmt.Errorf("query token err，code=%d，msg=%s，requestId=%s", response.Code, response.Msg, response.RequestId)
	}
	return &response, nil
}

// RefreshToken 刷新token
func (ac *AccessClient) RefreshToken(refreshToken string) (*JdAccessTokenResponse, error) {
	var response JdAccessTokenResponse
	link := fmt.Sprintf(RefreshToken, ac.AppKey, ac.AppSecret, refreshToken)

	if err := sdkHttp.Request(http.MethodGet, link, nil, &response); err != nil {
		return nil, fmt.Errorf("refresh access token failed, err=%+v", err)
	}
	if response.Code != 0 {
		return nil, fmt.Errorf("refresh token err，code=%d，msg=%s，requestId=%s", response.Code, response.Msg, response.RequestId)
	}
	return &response, nil
}
