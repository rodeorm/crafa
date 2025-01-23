package server

import (
	"money/internal/http/page"
	"net/http"
)

func (s *Server) forbidden(w http.ResponseWriter, r *http.Request) {
	session, _ := s.getSession(r)
	page.Execute("index", "forbidden", w, page.NewPage(page.WithSession(session)))
}
