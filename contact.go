package feishu

import (
	"net/http"
)

type ContactService struct {
	client *Client
}

type BatchGetIdOptions struct {
	Emails  []string `json:"emails,omitempty"`
	Mobiles []string `json:"mobiles,omitempty"`
}

type BatchGetIdQueryOptions struct {
	UserIdType string `url:"user_id_type"`
}

type User struct {
	UserId string `json:"user_id"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
}

type BatchUsers struct {
	CodeMsg
	Data struct {
		UserList []User `json:"user_list"`
	} `data`
}

func (s *ContactService) BatchGetId(idType string, opt *BatchGetIdOptions, options ...RequestOptionFunc) (*BatchUsers, *Response, error) {
	u := "contact/v3/users/batch_get_id"
	options = append(options, WithQuery(&BatchGetIdQueryOptions{UserIdType: idType}))

	req, err := s.client.NewServerRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	c := new(BatchUsers)
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, err
}
