package core

import (
	"context"
	"database/sql"

	"money/internal/crypt"
	"time"

	"github.com/golang-jwt/jwt"
)

type Session struct {
	ID int
	User

	LoginTime      time.Time
	LogoutTime     sql.NullTime
	LastActionTime sql.NullTime

	OTP string
}

type SessionStorager interface {
	UpdateSession(context.Context, *Session) error // Обновление сессии (совмещенное с проверкой)
}

// CodeSession кодирует сессию в строку c использованием JWT
// Для этого этой функции надо передать  данные сессии, ключ для кодирования, время жизни токена в миллисекундах
func CodeSession(s *Session, jwtKey string, tokenLiveTime time.Duration) (string, error) {
	key := []byte(jwtKey)
	c := crypt.CreateClaims(s.Login, s.ID, s.Role.ID, s.User.ID, tokenLiveTime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func GetSessionFromTkn(tknStr, jwtKey string) (*Session, error) {
	cl, err := crypt.GetClaims(tknStr, jwtKey)
	if err != nil {
		return nil, err
	}

	return &Session{ID: cl.SessionID, User: User{ID: cl.UserID, Login: cl.Login, Role: Role{ID: cl.RoleID}}}, nil
}
