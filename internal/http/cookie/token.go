package cookie

import (
	"net/http"
)

// NewCookieWithToken создает cookie с переданным токеном
func NewCookieWithToken(token string, maxAge int) *http.Cookie {
	// log.Printf("NewCookieWithToken %d", maxAge)
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

	/*
		log.Printf("Получили куки %s истекающие:%s\n ", c.Value, c.Expires)

		if c.Expires.Before(time.Now()) {
			return "", err
		}
	*/
	if err != nil {
		return "", err
	}
	return c.Value, nil
}
