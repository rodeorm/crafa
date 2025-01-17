package server

import (
	"money/internal/http/page"
	"net/http"
)

func (s *Server) regGet(w http.ResponseWriter, r *http.Request) {
	page.Execute("user", "reg", w, page.NewPage())
}
