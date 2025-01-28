package server

import (
	"log"
	"money/internal/core"
	"money/internal/http/page"
	"net/http"
)

func (s *Server) index(w http.ResponseWriter, r *http.Request) {

	session, err := s.getSession(r)
	if err != nil {
		page.Execute("index", "index", w, page.NewPage())
		return
	}
	p := page.NewPage(page.WithSession(session))
	switch session.User.Role.ID {
	case core.RoleGuest:
		page.Execute("index", "index", w, p)
	case core.RoleAdmin:
		page.Execute("admin", "index", w, p)
	case core.RoleReg:
		http.Redirect(w, r, "/user/send", http.StatusTemporaryRedirect)
	case core.RoleAuth:
		log.Println("HERE WE ARE")
		page.Execute("index", "auth", w, p)
	}
}
