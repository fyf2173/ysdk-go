package except

import (
	"context"
	"fmt"

	"github.com/fyf2173/ysdk-go/xhttp"
)

const (
	alertUrl = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s"
)

func NewAlert(key string, mode string) *RobotAlert {
	return &RobotAlert{Key: key, Mode: mode}
}

type RobotAlert struct {
	Mode string
	Key  string
}

func (al *RobotAlert) Url() string {
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
func (al *RobotAlert) SendText(ctx context.Context, content string) error {
	var req = MessageReq{
		Msgtype: "text",
		Text: &MessageContent{
			Content: content,
		},
	}

	if err := xhttp.Request(ctx, "POST", al.Url(), req, nil); err != nil {
		return err
	}
	return nil
}

// SendMarkdown 发送markdown通知
func (al *RobotAlert) SendMarkdown(ctx context.Context, content string) error {
	var req = MessageReq{
		Msgtype: "markdown",
		Markdown: &MessageContent{
			Content: content,
		},
	}

	if err := xhttp.Request(ctx, "POST", al.Url(), req, nil); err != nil {
		return err
	}
	return nil
}
