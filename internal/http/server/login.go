package server

import (
	"context"
	"net/http"

	"money/internal/core"
	"money/internal/http/cookie"
	"money/internal/http/page"
)

func (s *Server) loginPost(w http.ResponseWriter, r *http.Request) {
	user := &core.User{Login: r.FormValue("login"), Password: r.FormValue("password")}
	ctx := context.Background()

	err := s.storages.UserStorager.AuthUser(ctx, user)
	if err != nil {
		sign := make(map[string]string)
		sign["russ"] = "неправильное имя пользователя или пароль"
		sign["err"] = err.Error()
		pg := page.NewPage(page.WithSignals(sign))
		w.WriteHeader(http.StatusUnauthorized)
		page.Execute("index", "index", w, nil, pg)
		return
	}

	session, err := s.storages.SessionStorager.AddSession(ctx, user)
	if err != nil {
		sign := make(map[string]string)
		sign["russ"] = "внутренняя ошибка сервера"
		sign["err"] = err.Error()
		pg := page.NewPage(page.WithSignals(sign))
		w.WriteHeader(http.StatusUnauthorized)
		page.Execute("index", "index", w, nil, pg)
		return
	}
	cookie, err := cookie.NewCookieWithSession(session, s.cfg.JWTKey, s.cfg.TokeLiveTime)
	if err != nil {
		sign := make(map[string]string)
		sign["russ"] = "внутренняя ошибка сервера"
		sign["err"] = err.Error()
		pg := page.NewPage(page.WithSignals(sign))
		w.WriteHeader(http.StatusUnauthorized)
		page.Execute("index", "index", w, nil, pg)
		return
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

}
