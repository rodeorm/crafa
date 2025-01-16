package server

import (
	"net/http"
)

func (s *Server) forbidden(w http.ResponseWriter, r *http.Request) {
	// currentInformation, _ := h.getSessionInformation(r)
	// pages.ExecuteHTML("index", "forbidden", w, *currentInformation)
}
