package ui

import "net/http"

func (h Handler) regGet(w http.ResponseWriter, r *http.Request) {
	ExecuteHTML("user", "reg", w, sessionInformation{})
}
