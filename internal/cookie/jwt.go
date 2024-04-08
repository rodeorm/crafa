package cookie

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

// JWT ключ, используемый для создания подписи (TODO: вынести в конфиг)
var jwtKey = []byte("top_secret_key")

func PutTokenToCookie(c *Claims) *http.Cookie {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, _ := token.SignedString(jwtKey)

	return &http.Cookie{
		Name:   "token",
		Value:  tokenString,
		MaxAge: 3600,
		Path:   "/",
	}

}

func RemoveTokenFromCookie() *http.Cookie {
	return &http.Cookie{
		Name:   "token",
		MaxAge: -100,
	}
}
