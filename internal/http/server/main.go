package server

import (
	"money/internal/core"
	"money/internal/http/page"
	"net/http"
)

func (s *Server) main(w http.ResponseWriter, r *http.Request) {

	session, err := s.getSession(r)
	if err != nil {
		page.Execute("index", "index", w, page.NewPage())
		return
	}

	switch session.User.Role.ID {
	case core.RoleAdmin:
		p := page.NewPage(page.WithSession(session))
		page.Execute("admin", "index", w, p)
	case core.RoleReg:
		http.Redirect(w, r, "/user/send", http.StatusTemporaryRedirect)
	case core.RoleAuth:
		p := page.NewPage(page.WithSession(session))
		page.Execute("main", "auth", w, p)
	}
}
