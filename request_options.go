package feishu

import (
	"context"
	"github.com/hashicorp/go-retryablehttp"
)

// RequestOptionFunc can be passed to all API requests to customize the API request.
type RequestOptionFunc func(*retryablehttp.Request) error

// WithContext runs the request with the provided context
func WithContext(ctx context.Context) RequestOptionFunc {
	return func(req *retryablehttp.Request) error {
		*req = *req.WithContext(ctx)
		return nil
	}
}

// WithToken takes a token which is then used when making this one request.
func WithToken(token string) RequestOptionFunc {
	return func(req *retryablehttp.Request) error {
		req.Header.Set("Authorization", "Bearer "+token)
		return nil
	}
}
