package ui

import (
	"log"
	"net/http"
	"time"

	"money/internal/cookie"
)

func (h Handler) logOutPost(w http.ResponseWriter, r *http.Request) {
	claims, err := cookie.GetClaimsFromCookie(r)
	if err != nil {
		log.Println("logOutPost", err)
		ExecuteHTML("index", "index", w, sessionInformation{})
		return
	}
	session, err := h.Auth.SelectSession(claims.UserID, claims.SessionID)
	if err != nil {
		log.Println("logOutPost", err)
		ExecuteHTML("index", "index", w, sessionInformation{})
		return
	}
	session.LogOutTime.Time = time.Now()
	session.LogOutTime.Valid = true
	h.Auth.UpdateSession(session)

	http.SetCookie(w, cookie.RemoveTokenFromCookie())
	ExecuteHTML("index", "logout", w, sessionInformation{})
}
