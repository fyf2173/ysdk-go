package auth

type AccessClient struct {
	AppKey    string
	AppSecret string
}

func NewAccessClient(appKey, appSecret string) *AccessClient {
	return &AccessClient{AppKey: appKey, AppSecret: appSecret}
}
