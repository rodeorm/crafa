package crypt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	RoleID    int    `json:"roleid"`
	UserID    int    `json:"userid"`
	SessionID int    `json:"sessionid"`
	Login     string `json:"login"`
	jwt.StandardClaims
}

// createClaims создает jwt-claims. Время жизни токена в миллисекундах!
func CreateClaims(login string, sessionID, roleID, userID int, liveTime time.Duration) *Claims {
	return &Claims{
		SessionID: sessionID,
		Login:     login,
		UserID:    userID,
		RoleID:    roleID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(liveTime * time.Millisecond).Unix(),
		},
	}
}

func GetClaims(tknStr string) (*Claims, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !tkn.Valid {
		return nil, fmt.Errorf("невалидный токен %v", err)
	}
	return claims, nil
}
