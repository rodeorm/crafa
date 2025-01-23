package server

import (
	"context"
	"money/internal/core"
	"money/internal/http/page"
	"net/http"
)

func (s *Server) verifyPost(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login") // из скрытого поля
	email := r.FormValue("email") // из скрытого поля
	otp := r.FormValue("otp")

	usr := &core.User{Login: login, Email: email}
	err := s.stgs.UserStorager.AdvAuthUser(context.TODO(), usr, otp)
	if err != nil {
		sign := make(map[string]string)
		sign["russ"] = "неправильный код подтверждения"
		sign["err"] = err.Error()
		pg := page.NewPage(page.WithSignals(sign))
		w.WriteHeader(http.StatusUnauthorized)
		page.Execute("index", "verify", w, pg)
		return
	}
}
