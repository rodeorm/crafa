package cookie

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Login     string `json:"login"`
	RoleID    int    `json:"roleid"`
	UserID    int    `json:"userid"`
	RoleName  string `json:"rolename"`
	SessionID int    `json:"sessionid"`
	jwt.StandardClaims
}

func CreateClaims(login string, sessionID, roleID, userID int, roleName string) *Claims {
	expirationTime := time.Now().Add(5000 * time.Minute)

	return &Claims{
		SessionID: sessionID,
		Login:     login,
		UserID:    userID,
		RoleID:    roleID,
		RoleName:  roleName,
		StandardClaims: jwt.StandardClaims{
			// в JWT время жизни в Unix миллисекундах
			ExpiresAt: expirationTime.Unix(),
		},
	}
}

func GetClaimsFromCookie(r *http.Request) (*Claims, error) {
	c, err := r.Cookie("token")
	if err != nil {
		return nil, err
	}
	tknStr := c.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		log.Println("GetClaimsFromCoockie", "ошибка парсинга", r.URL, err)
		return nil, err
	}
	if !tkn.Valid {
		log.Println("GetClaimsFromCoockie", "невалидный токен", r.URL)
		return nil, fmt.Errorf("невалидный токен")
	}
	return claims, nil
}
