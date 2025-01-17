package server

import (
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
	case core.Guest:
		page.Execute("index", "index", w, p)
	case core.Admin:
		page.Execute("admin", "accountList", w, p)
	case core.Reg:
		page.Execute("index", "indexUnAuth", w, p)
	case core.Auth:
		page.Execute("index", "indexAuth", w, p)
	}
}
