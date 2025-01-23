package server

import (
	"context"
	"money/internal/core"
	"money/internal/http/page"
	"net/http"
)

func (s *Server) loginPost(w http.ResponseWriter, r *http.Request) {
	user := &core.User{Login: r.FormValue("login"), Password: r.FormValue("password")}
	ctx := context.Background()

	err := s.stgs.UserStorager.BaseAuthUser(ctx, user)
	if err != nil {
		sign := make(map[string]string)
		sign["russ"] = "неправильное имя пользователя или пароль"
		sign["err"] = err.Error()
		pg := page.NewPage(page.WithSignals(sign))
		w.WriteHeader(http.StatusUnauthorized)
		page.Execute("index", "index", w, pg)
		return
	}

	at := make(map[string]any)
	at["login"] = user.Login
	at["email"] = user.Password

	page.Execute("index", "verify", w, page.NewPage(page.WithAttrs(at)))
}
