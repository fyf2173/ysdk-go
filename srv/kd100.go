package srv

const (
	KD100SubscribeURL       = "https://poll.kuaidi100.com/poll"
	KD100RequestContentType = "application/x-www-form-urlencoded"
)

// KD100RequestBody 订阅请求体
type KD100RequestBody struct {
	Schema string `json:"schema"`
	Param  string `json:"param"` // KD100ResponseBody => json encode
}

type KD100ResponseBody struct {
	Result     bool   `json:"result"`
	ReturnCode string `json:"returnCode"`
	Message    string `json:"message"`
}

type KD100Param struct {
	Company    string          `json:"company"`        // 订阅的快递公司的编码，一律用小写字母
	Number     string          `json:"number"`         // 单号
	From       string          `json:"from,omitempty"` // 出发地城市，省-市-区
	To         string          `json:"to,omitempty"`   // 目的地城市，省-市-区
	Key        string          `json:"key"`            // 授权码
	Parameters KD100Parameters `json:"parameters"`     // 附加参数信息
}

type KD100Parameters struct {
	Callbackurl        string `json:"callbackurl"`                  // 回调接口的地址
	Salt               string `json:"salt,omitempty"`               // 签名用随机字符串
	Resultv2           string `json:"resultv2,omitempty"`           // 添加此字段表示打开行政区域解析功能, 例：1
	AutoCom            string `json:"autoCom,omitempty"`            // 表示开始智能判断单号所属公司的功能，例：1
	InterCom           string `json:"interCom,omitempty"`           // 表示开启国际版，例：1
	DepartureCountry   string `json:"departureCountry,omitempty"`   // 出发国家编码
	DepartureCom       string `json:"departureCom,omitempty"`       // 出发的快递公司的编码
	DestinationCountry string `json:"destinationCountry,omitempty"` // 目的国家编码
	DestinationCom     string `json:"destinationCom,omitempty"`     // 目的的快递公司的编码
	Phone              string `json:"phone"`                        // 顺丰单号必填，其他快递公司选填
}

// KD100CallbackRequestBody 回调请求体
type KD100CallbackRequestBody struct {
}
