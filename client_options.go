package feishu

import (
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
)

// ClientOptionFunc can be used to customize a new GitLab API client.
type ClientOptionFunc func(*Client) error

// WithBaseURL sets the base URL for API requests to a custom endpoint.
func WithBaseURL(urlStr string) ClientOptionFunc {
	return func(c *Client) error {
		return c.setBaseURL(urlStr)
	}
}

// WithCustomBackoff can be used to configure a custom backoff policy.
func WithCustomBackoff(backoff retryablehttp.Backoff) ClientOptionFunc {
	return func(c *Client) error {
		c.client.Backoff = backoff
		return nil
	}
}

// WithCustomLeveledLogger can be used to configure a custom retryablehttp
// leveled logger.
func WithCustomLeveledLogger(leveledLogger retryablehttp.LeveledLogger) ClientOptionFunc {
	return func(c *Client) error {
		c.client.Logger = leveledLogger
		return nil
	}
}

// WithCustomLimiter injects a custom rate limiter to the client.
func WithCustomLimiter(limiter RateLimiter) ClientOptionFunc {
	return func(c *Client) error {
		c.configureLimiterOnce.Do(func() {
			c.limiter = limiter
		})
		return nil
	}
}

// WithCustomLogger can be used to configure a custom retryablehttp logger.
func WithCustomLogger(logger retryablehttp.Logger) ClientOptionFunc {
	return func(c *Client) error {
		c.client.Logger = logger
		return nil
	}
}

// WithCustomRetry can be used to configure a custom retry policy.
func WithCustomRetry(checkRetry retryablehttp.CheckRetry) ClientOptionFunc {
	return func(c *Client) error {
		c.client.CheckRetry = checkRetry
		return nil
	}
}

// WithHTTPClient can be used to configure a custom HTTP client.
func WithHTTPClient(httpClient *http.Client) ClientOptionFunc {
	return func(c *Client) error {
		c.client.HTTPClient = httpClient
		return nil
	}
}

// WithRequestLogHook can be used to configure a custom request log hook.
func WithRequestLogHook(hook retryablehttp.RequestLogHook) ClientOptionFunc {
	return func(c *Client) error {
		c.client.RequestLogHook = hook
		return nil
	}
}

// WithResponseLogHook can be used to configure a custom response log hook.
func WithResponseLogHook(hook retryablehttp.ResponseLogHook) ClientOptionFunc {
	return func(c *Client) error {
		c.client.ResponseLogHook = hook
		return nil
	}
}

// WithoutRetries disables the default retry logic.
func WithoutRetries() ClientOptionFunc {
	return func(c *Client) error {
		c.disableRetries = true
		return nil
	}
}

func WithTenantAccessTokenInternal(appId, appSecret string) ClientOptionFunc {
	return func(c *Client) error {
		refreshFunc := func() (*TenantAccessToken, error) {
			opt := &GetAccessTokenOptions{
				AppId:     appId,
				AppSecret: appSecret,
			}
			if t, _, err := c.Auth.GetTenantAccessTokenInternal(opt); err != nil {
				return nil, err
			} else {
				return t, nil
			}
		}
		c.accessTokenManager = NewAccessTokenManager(appId, _tenantAccessTokenInternal, refreshFunc, WithLocalCache())
		return nil
	}
}

func WithTenantAccessTokenInternalRedis(appId, appSecret, addr, password string) ClientOptionFunc {
	return func(c *Client) error {
		refreshFunc := func() (*TenantAccessToken, error) {
			opt := &GetAccessTokenOptions{
				AppId:     appId,
				AppSecret: appSecret,
			}
			if t, _, err := c.Auth.GetTenantAccessTokenInternal(opt); err != nil {
				return nil, err
			} else {
				return t, nil
			}
		}
		c.accessTokenManager = NewAccessTokenManager(appId, _tenantAccessTokenInternal, refreshFunc, WithRedisCache(addr, password))
		return nil
	}
}
