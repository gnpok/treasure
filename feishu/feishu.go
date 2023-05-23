package feishu

import (
	"github.com/beego/beego/v2/adapter/httplib"
)

type H map[string]interface{}

type FeiShu struct {
	Domain    string
	MsgPrefix string
}

func NewFeiShu(domain string) *FeiShu {
	return &FeiShu{
		Domain: domain,
	}
}

// SetDomain 修改飞书通知地址
func (f *FeiShu) SetDomain(domain string) {
	if domain == "" {
		return
	}
	f.Domain = domain
}

// BaseNotify 飞书基础消息通知
func (f *FeiShu) BaseNotify(msg string) (err error) {
	payload := H{
		"msg_type": "text",
		"content": H{
			"text": msg,
		},
	}
	req := httplib.Post(f.Domain)
	req.JSONBody(payload)
	req.Header("Content-Type", "application/json")
	_, err = req.String()
	return
}

// ImportanceNotify 通知重要的信息
func (f *FeiShu) ImportanceNotify(title, content string) (err error) {
	payload := H{
		"msg_type": "post",
		"content": H{
			"post": H{
				"zh_cn": H{
					"title": title,
					"content": [][]H{
						{H{
							"tag":  "text",
							"text": content,
						}},
					},
				},
			},
		},
	}
	req := httplib.Post(f.Domain)
	req.JSONBody(payload)
	req.Header("Content-Type", "application/json")
	_, err = req.String()
	return
}
