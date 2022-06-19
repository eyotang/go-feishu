package feishu

import "time"

const (
	_appAccessToken         = "app_access_token"
	_appAccessTokenInternal = "app_access_token_internal"

	_tenantAccessToken         = "tenant_access_token"
	_tenantAccessTokenInternal = "tenant_access_token_internal"
)

type AccessTokenManager interface {
	GetAccessToken() (err error, accessToken string)
}

type TokenRefreshFunc func() (string, error)

type accessTokenManagerService struct {
	refreshFunc TokenRefreshFunc
	Id          string
	tokenType   string
	Cache       TokenCache
}

func NewAccessTokenManager(appId string, tokenType string, refreshFunc TokenRefreshFunc, options ...CacheOptionFunc) (atms *accessTokenManagerService) {
	atms = &accessTokenManagerService{
		Id:          appId,
		tokenType:   tokenType,
		refreshFunc: refreshFunc,
	}
	for _, fn := range options {
		if fn == nil {
			continue
		}
		fn(atms)
	}
	return
}

func (s *accessTokenManagerService) TokenKey() string {
	return s.tokenType + ":" + s.Id
}

func (s *accessTokenManagerService) GetAccessToken() (err error, accessToken string) {
	tokenKey := s.TokenKey()

	// 未过期，直接使用
	if value, ok := s.Cache.Get(tokenKey); ok {
		accessToken = value.(string)
		return
	}

	// 上锁
	s.Cache.Lock()
	defer s.Cache.Unlock()

	// 其他协程获取并保存了，直接使用
	if value, ok := s.Cache.Get(tokenKey); ok {
		accessToken = value.(string)
		return
	}

	if accessToken, err = s.refreshFunc(); err != nil {
		return
	}

	// 保存token
	expire := 2*time.Hour - defaultCleanup - time.Second
	s.Cache.Set(tokenKey, accessToken, expire)
	return
}
