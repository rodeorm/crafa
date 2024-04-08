package ui

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"money/internal/cookie"
	"money/internal/core"
)

func (h Handler) getSessionInformation(r *http.Request) (*sessionInformation, error) {
	claims, err := cookie.GetClaimsFromCookie(r)
	if err != nil {
		log.Println("getCurrentUserInformation", err)
		return &sessionInformation{}, fmt.Errorf("пользователь не найден в куках")
	}

	session, err := h.Auth.SelectActiveSession(claims.UserID, claims.SessionID)
	if err != nil {
		log.Println("getCurrentUserInformation", err)
		return &sessionInformation{}, fmt.Errorf("для пользователя нет активной сессии")
	}
	session.LastActionTime.Time = time.Now()
	session.LastActionTime.Valid = true
	h.Auth.UpdateSession(session)

	return &sessionInformation{User: core.User{ID: claims.UserID, Login: claims.Login, Role: core.Role{ID: claims.RoleID, Name: claims.RoleName}}}, nil
}
