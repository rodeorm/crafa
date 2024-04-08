package ui

import (
	"net/http"
)

func (h Handler) forbidden(w http.ResponseWriter, r *http.Request) {
	currentInformation, _ := h.getSessionInformation(r)
	ExecuteHTML("index", "forbidden", w, *currentInformation)
}
