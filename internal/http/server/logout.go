package server

import (
	"net/http"
	"time"
)

func (s *Server) logOut(w http.ResponseWriter, r *http.Request) {
	session, _ := s.getSession(r)
	session.LogoutTime.Time = time.Now()
	session.LogoutTime.Valid = true

	ck := &http.Cookie{
		Name:  "token",
		Value: "",
		Path:  "/",
	}
	http.SetCookie(w, ck)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
