package cookie

import (
	"money/internal/core"
	"net/http"
	"time"
)

// NewCookieWithSession возвращает http.Cookie с данными сессии в jwt-токене
func NewCookieWithSession(s *core.Session, key string, liveTime time.Duration) (*http.Cookie, error) {
	token, err := core.CodeSession(s, key, liveTime)
	if err != nil {
		return nil, err
	}
	liveTimeInSecond := int(liveTime / 1000)
	return NewCookieWithToken(token, liveTimeInSecond), nil
}
