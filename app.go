package feishu

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AppService struct {
	client *Client
}

type AppCardMessageOption struct {
	MsgType   string        `json:"msg_type"`
	ReceiveID string        `json:"receive_id"`
	Card      AppCardOption `json:"-"`       // 屏蔽差异
	Content   string        `json:"content"` // 飞书应用: content是字符串，需要将json object先marshal为字符串。飞书机器人: card是json object。
}

type AppCardOption struct {
	Config   CardConfigOption `json:"config"`
	Header   HeadOption       `json:"header"`
	Elements []interface{}    `json:"elements"`
}

type AppCardMessageQueryOptions struct {
	ReceiveIdType string `url:"receive_id_type"`
}

func (s *AppService) SendAppCardMessage(receiveIDType string, opt *AppCardMessageOption, options ...RequestOptionFunc) (*ErrorMessage, *Response, error) {
	u := fmt.Sprintf("im/v1/messages")
	options = append(options, WithQuery(&AppCardMessageQueryOptions{ReceiveIdType: receiveIDType}))

	// 尊重意愿，可以调用者在外面自己marshal。
	if opt != nil && len(opt.Content) == 0 {
		if buf, err := json.Marshal(opt.Card); err != nil {
			return nil, nil, err
		} else {
			opt.Content = string(buf)
		}
	}
	req, err := s.client.NewServerRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	c := new(ErrorMessage)
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, err
}
