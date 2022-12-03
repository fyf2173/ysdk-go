package except

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/fyf2173/ysdk-go/xhttp"
)

const (
	alertUrl = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s"
)

func NewAlert(key string, mode string) *Alert {
	return &Alert{Key: key, Mode: mode}
}

type Alert struct {
	Mode string
	Key  string
}

func (al *Alert) Url() string {
	return fmt.Sprintf(alertUrl, al.Key)
}

type MessageReq struct {
	Msgtype  string          `json:"msgtype"`
	Text     *MessageContent `json:"text,omitempty"`
	Markdown *MessageContent `json:"markdown,omitempty"`
}

type MessageContent struct {
	Content string `json:"content"`
}

// SendText 发送文本通知
func (al *Alert) SendText(content string) error {
	var req = MessageReq{
		Msgtype: "text",
		Text: &MessageContent{
			Content: content,
		},
	}

	var resp interface{}
	if err := xhttp.Request("POST", al.Url(), req, resp); err != nil {
		return err
	}
	return nil
}

type MarkdownAlertTopic struct {
	Mode         string
	Title        string
	EnterpriseId int64
	Desc         string
	RequestId    string
	AlertAt      time.Time
	LineAt       string
	Params       string
}

type Option func(topic *MarkdownAlertTopic)

func WithEnterpriseId(enterpriseId int64) Option {
	return func(topic *MarkdownAlertTopic) {
		topic.EnterpriseId = enterpriseId
	}
}

func WithParams(params interface{}) Option {
	return func(topic *MarkdownAlertTopic) {
		b, _ := json.Marshal(params)
		topic.Params = string(b)
	}
}

func NewMarkdownAlertTopic(title, desc string, opts ...Option) *MarkdownAlertTopic {
	mat := &MarkdownAlertTopic{
		Title:   title,
		Desc:    desc,
		AlertAt: time.Now(),
	}
	for _, opt := range opts {
		opt(mat)
	}
	return mat
}

func (mat *MarkdownAlertTopic) String() string {
	var md string
	md += fmt.Sprintf("系统触发业务告警【<font color=\"warning\">%s</font>】，请及时排查。\n", mat.Title)
	md += fmt.Sprintf(">环境: <font color=\"comment\">%s</font> \n", mat.Mode)
	if mat.EnterpriseId > 0 {
		md += fmt.Sprintf(">企业ID: <font color=\"comment\">%d</font> \n", mat.EnterpriseId)
	}
	md += fmt.Sprintf(">描述: <font color=\"comment\">%s</font> \n", mat.Desc)
	md += fmt.Sprintf(">请求ID: <font color=\"comment\">%s</font> \n", mat.RequestId)
	md += fmt.Sprintf(">时间: <font color=\"comment\">%s</font> \n", mat.AlertAt.Format("2006-01-02 15:04:05"))
	if mat.Params != "" {
		md += fmt.Sprintf(">参数: <font color=\"comment\">%s</font> \n", mat.Params)
	}
	return md
}

// SendMarkdown 发送markdown通知
func (al *Alert) SendMarkdown(content string) error {
	var req = MessageReq{
		Msgtype: "markdown",
		Markdown: &MessageContent{
			Content: content,
		},
	}

	var resp interface{}
	if err := xhttp.Request("POST", al.Url(), req, resp); err != nil {
		return err
	}
	return nil
}
