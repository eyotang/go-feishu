package feishu

import (
	"net/http"
)

type AuthService struct {
	client *Client
}

type GetAccessTokenOptions struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type AppAccessTokenInternal struct {
	CodeMsg
	AccessToken string `json:"app_access_token"`
	Expire      int    `json:"expire"`
}

func (s *AuthService) GetAppAccessTokenInternal(opt *GetAccessTokenOptions, options ...RequestOptionFunc) (*AppAccessTokenInternal, *Response, error) {
	u := "auth/v3/app_access_token/internal"

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	c := new(AppAccessTokenInternal)
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, err
}

type AppAccessToken struct {
	AppAccessTokenInternal
	AppTicket string `json:"app_ticket"`
}

func (s *AuthService) GetAppAccessToken(opt *GetAccessTokenOptions, options ...RequestOptionFunc) (*AppAccessToken, *Response, error) {
	u := "auth/v3/app_access_token"

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	c := new(AppAccessToken)
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, err
}

type TenantAccessToken struct {
	CodeMsg
	AccessToken string `json:"tenant_access_token"`
	Expire      int    `json:"expire"`
}

func (s *AuthService) GetTenantAccessTokenInternal(opt *GetAccessTokenOptions, options ...RequestOptionFunc) (*TenantAccessToken, *Response, error) {
	u := "auth/v3/tenant_access_token/internal"

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	c := new(TenantAccessToken)
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, err
}

type TenantAccessTokenOptions struct {
	AppAccessToken string `json:"app_access_token"`
	TenantKey      string `json:"tenant_key"`
}

func (s *AuthService) GetTenantAccessToken(opt *GetAccessTokenOptions, options ...RequestOptionFunc) (*TenantAccessToken, *Response, error) {
	u := "auth/v3/tenant_access_token"

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	c := new(TenantAccessToken)
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, err
}
