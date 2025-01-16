package crypt

import (
	"money/internal/core"
	"time"

	"github.com/golang-jwt/jwt"
)

// CodeSession кодирует сессию в строку c использованием JWT
// Для этого этой функции надо передать  данные сессии, ключ для кодирования, время жизни токена в миллисекундах
func CodeSession(s *core.Session, jwtKey string, tokenLiveTime time.Duration) (string, error) {
	key := []byte(jwtKey)
	c := createClaims(s.Login, s.ID, s.Role.ID, s.User.ID, tokenLiveTime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
