package cookie

import (
	"net/http"
)

func NewCookieWithToken(token string, maxAge int) *http.Cookie {
	return &http.Cookie{
		Name:   "token",
		Value:  token,
		MaxAge: maxAge,
		Path:   "/",
	}

}

func RemoveTokenFromCookie() *http.Cookie {
	return &http.Cookie{
		Name:   "token",
		MaxAge: -100,
	}
}

func GetTokenFromRequest(r *http.Request) (string, error) {
	c, err := r.Cookie("token")
	if err != nil {
		return "", err
	}
	return c.Value, nil
}
