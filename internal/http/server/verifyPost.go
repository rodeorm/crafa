package server

import (
	"context"
	"log"
	"money/internal/core"
	"money/internal/http/cookie"
	"money/internal/http/page"
	"net/http"
	"strconv"
)

func (s *Server) verifyPost(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
		return
	}

	otp := r.FormValue("otp")
	usr := &core.User{ID: id}

	session, err := s.stgs.UserStorager.AdvAuthUser(context.TODO(), usr, otp, s.cfg.OTPLiveTime)
	if err != nil {
		sign := make(map[string]string)
		sign["russ"] = "Неправильный код подтверждения"
		sign["err"] = err.Error()
		pg := page.NewPage(page.WithSignals(sign))
		w.WriteHeader(http.StatusUnauthorized)
		page.Execute("user", "verify", w, pg)
		return
	}

	log.Println("verifyPost", session)

	ck, err := cookie.NewCookieWithSession(session, s.cfg.JWTKey, s.cfg.TokenLiveTime)
	if err != nil {
		sign := make(map[string]string)
		sign["russ"] = "Ошибка при аутентификации"
		sign["err"] = err.Error()
		at := make(map[string]any)
		at["User"] = usr
		pg := page.NewPage(page.WithSignals(sign), page.WithAttrs(at))
		w.WriteHeader(http.StatusUnauthorized)
		page.Execute("user", "verify", w, pg)
		return
	}

	http.SetCookie(w, ck)
	http.Redirect(w, r, "/main", http.StatusTemporaryRedirect)
}
