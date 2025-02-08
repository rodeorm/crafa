package server

import (
	"context"
	"log"
	"money/internal/core"
	"money/internal/http/cookie"
	"money/internal/http/page"
	"net/http"
)

func (s *Server) regPost(w http.ResponseWriter, r *http.Request) {

	// Получаем данные из формы
	user := core.User{
		Login:      r.FormValue("login"),
		Password:   r.FormValue("password"),
		Name:       r.FormValue("name"),
		PatronName: r.FormValue("patronname"),
		FamilyName: r.FormValue("familyname"),
		Email:      r.FormValue("email"),
		Phone:      r.FormValue("phonenumber"),
	}
	// Регистрируем пользователя. Получаем идентификатор пользователя и идентификатор сессии
	session, err := s.stgs.RegUser(context.TODO(), &user, s.cfg.Domain)
	if err != nil {
		sign := make(map[string]string)
		sign["Russ"] = "Ошибка при регистрации"
		sign["Err"] = err.Error()
		at := make(map[string]any)
		at["User"] = user
		pg := page.NewPage(page.WithSignals(sign), page.WithAttrs(at))
		log.Println(pg)
		w.WriteHeader(http.StatusUnauthorized)
		page.Execute("user", "reg", w, pg)
		return
	}
	// Создаем jwt-токен и сохраняем его в куках
	ck, err := cookie.NewCookieWithSession(session, s.cfg.JWTKey, s.cfg.TokenLiveTime)
	if err != nil {
		sign := make(map[string]string)
		sign["russ"] = "Ошибка при регистрации"
		sign["err"] = err.Error()

		at := make(map[string]any)
		at["User"] = user
		pg := page.NewPage(page.WithSignals(sign), page.WithAttrs(at))
		w.WriteHeader(http.StatusUnauthorized)
		page.Execute("user", "reg", w, pg)
		return
	}

	http.SetCookie(w, ck)

	http.Redirect(w, r, "/user/send", http.StatusTemporaryRedirect)
}
