package server

import (
	"context"
	"net/http"
	"time"

	"money/internal/http/cookie"
	"money/internal/http/page"
)

func (s *Server) logOutPost(w http.ResponseWriter, r *http.Request) {
	session, err := s.getSession(r)

	if err != nil {
		page.Execute("index", "index", w, session, nil)
		return
	}

	session.LogoutTime.Time = time.Now()
	session.LogoutTime.Valid = true

	s.storages.SessionStorager.UpdateSession(context.TODO(), session)
	http.SetCookie(w, cookie.RemoveTokenFromCookie())
	page.Execute("index", "logout", w, nil, nil)
}
