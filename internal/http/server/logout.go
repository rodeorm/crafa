package server

import (
	"context"
	"money/internal/http/cookie"
	"money/internal/http/page"
	"net/http"
	"time"
)

func (s *Server) logOutPost(w http.ResponseWriter, r *http.Request) {
	session, err := s.getSession(r)

	p := page.NewPage(page.WithSession(session))

	if err != nil {
		page.Execute("index", "index", w, p)
		return
	}

	session.LogoutTime.Time = time.Now()
	session.LogoutTime.Valid = true

	s.storages.SessionStorager.UpdateSession(context.TODO(), session)
	http.SetCookie(w, cookie.RemoveTokenFromCookie())
	page.Execute("index", "logout", w, page.NewPage())
}
