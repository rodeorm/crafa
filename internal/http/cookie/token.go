package cookie

import (
	"net/http"
)

// NewCookieWithToken создает cookie с переданным токеном
func NewCookieWithToken(token string, maxAge int) *http.Cookie {
	return &http.Cookie{
		Name:   "token",
		Value:  token,
		MaxAge: maxAge,
		Path:   "/",
	}

}

// RemoveTokenFromCookie удаляет cookie c именем "token"
func RemoveTokenFromCookie() *http.Cookie {
	return &http.Cookie{
		Name:   "token",
		MaxAge: -100,
	}
}

// GetTokenFromRequest возвращает токен из request
func GetTokenFromRequest(r *http.Request) (string, error) {
	c, err := r.Cookie("token")
	if err != nil {
		return "", err
	}
	return c.Value, nil
}
