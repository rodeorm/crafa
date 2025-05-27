package cookie

import (
	"money/internal/core"
	"net/http"
)

// NewCookieWithSession возвращает http.Cookie с данными сессии в jwt-токене
func (c CookieManager) NewCookieWithSession(s *core.Session) (*http.Cookie, error) {
	token, err := core.CodeSession(s, c.key, c.liveTime)
	if err != nil {
		return nil, err
	}
	liveTimeInSecond := int(c.liveTime / 1000)
	return NewCookieWithToken(token, liveTimeInSecond), nil
}
