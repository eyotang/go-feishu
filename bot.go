package feishu

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type BotService struct {
	client *Client
}

type BotCardMessageOption struct {
	Timestamp string        `json:"timestamp,omitempty"`
	Sign      string        `json:"sign,omitempty"`
	MsgType   string        `json:"msg_type"`
	Card      BotCardOption `json:"card"`
}

type BotCardOption struct {
	Config   CardConfigOption `json:"config"`
	Header   HeadOption       `json:"header"`
	Elements []interface{}    `json:"elements"`
}

type CardConfigOption struct {
	WideScreenMode bool `json:"wide_screen_mode"`
	EnableForward  bool `json:"enable_forward"`
}

type HeadOption struct {
	Title    TitleOption `json:"title"`
	Template string      `json:"template"`
}

type TitleOption struct {
	Tag     string `json:"tag"`
	Content string `json:"content"`
}

type BotResponse struct {
	Extra         interface{} `json:"Extra"`
	StatusCode    int         `json:"StatusCode"`
	StatusMessage string      `json:"StatusMessage"`
}

func (s *BotService) SendBotCardMessage(botKey, secret string, opt *BotCardMessageOption, options ...RequestOptionFunc) (*ErrorMessage, *Response, error) {
	u := fmt.Sprintf("bot/v2/hook/%s", botKey)

	if len(secret) > 0 {
		timestamp := time.Now().Unix()
		opt.Timestamp = strconv.FormatInt(timestamp, 10)
		if sign, err := GenSign(secret, timestamp); err != nil {
			return nil, nil, err
		} else {
			opt.Sign = sign
		}
	}

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
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

func GenSign(secret string, timestamp int64) (string, error) {
	//timestamp + key 做sha256, 再进行base64 encode
	stringToSign := fmt.Sprintf("%v", timestamp) + "\n" + secret
	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	_, err := h.Write(data)
	if err != nil {
		return "", err
	}
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature, nil
}

type BotMessageOption struct {
	Timestamp string      `json:"timestamp,omitempty"`
	Sign      string      `json:"sign,omitempty"`
	MsgType   string      `json:"msg_type"`
	Content   interface{} `json:"content"`
}

func (s *BotService) SendBotMessage(botKey, secret string, opt *BotMessageOption, options ...RequestOptionFunc) (*BotResponse, *Response, error) {
	u := fmt.Sprintf("bot/v2/hook/%s", botKey)

	if len(secret) > 0 {
		timestamp := time.Now().Unix()
		opt.Timestamp = strconv.FormatInt(timestamp, 10)
		if sign, err := GenSign(secret, timestamp); err != nil {
			return nil, nil, err
		} else {
			opt.Sign = sign
		}
	}

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	c := new(BotResponse)
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, err
}
