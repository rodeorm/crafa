package cookie

import (
	"money/internal/core"
	"money/internal/crypt"
	"net/http"
	"time"
)

func NewCookieWithSession(s *core.Session, key string, liveTime time.Duration) (*http.Cookie, error) {
	token, err := crypt.CodeSession(s, key, liveTime)
	if err != nil {
		return nil, err
	}
	liveTimeInSecond := int(liveTime / 1000)
	return NewCookieWithToken(token, liveTimeInSecond), nil
}
