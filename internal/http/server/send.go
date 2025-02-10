package server

import (
	"context"
	"money/internal/http/page"
	"net/http"
)

func (s *Server) send(w http.ResponseWriter, r *http.Request) {
	session, err := s.getSession(r)
	if err != nil {
		http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
		return
	}

	s.stgs.UserStorager.SelectUser(context.TODO(), &session.User)
	page.Execute("user", "send", w, page.NewPage(page.WithSession(session)))
}
