package server

import (
	"log"
	"net/http"
)

func (s *Server) forbidden(w http.ResponseWriter, r *http.Request) {
	log.Println("forbidden")
	// currentInformation, _ := h.getSessionInformation(r)
	// pages.ExecuteHTML("index", "forbidden", w, *currentInformation)
}
