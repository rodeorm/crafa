package server

import (
	"money/internal/core"
	"money/internal/http/page"
	"net/http"
)

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	session, err := s.getSession(r)

	if err != nil {
		page.Execute("index", "index", w, session, nil)
		return
	}

	switch session.User.Role.ID {
	case core.Guest:
		page.Execute("index", "index", w, session, nil)
	case core.Admin:
		page.Execute("admin", "accountList", w, nil, nil)
	case core.Reg:
		page.Execute("index", "indexUnAuth", w, nil, nil)
	case core.Auth:
		page.Execute("index", "indexAuth", w, nil, nil)
	}

}
