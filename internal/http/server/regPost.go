package server

import (
	"context"
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
		Patronymic: r.FormValue("patronname"),
		Surname:    r.FormValue("familyname"),
		Email:      r.FormValue("email"),
		Phone:      r.FormValue("phonenumber"),
		Role:       core.Role{ID: core.Reg},
	}
	// Регистрируем пользователя. Получаем идентификатор пользователя и идентификатор сессии
	session, err := s.stgs.RegUser(context.TODO(), &user)
	if err != nil {
		sign := make(map[string]string)
		sign["russ"] = "Ошибка при регистрации"
		sign["err"] = err.Error()

		at := make(map[string]any)
		at["user"] = user
		pg := page.NewPage(page.WithSignals(sign), page.WithAttrs(at))
		w.WriteHeader(http.StatusUnauthorized)
		page.Execute("index", "index", w, pg)
		return
	}
	// Создаем jwt-токен и сохраняем его в куках
	ck, err := cookie.NewCookieWithSession(session, s.cfg.JWTKey, s.cfg.TokeLiveTime)
	if err != nil {
		sign := make(map[string]string)
		sign["russ"] = "Ошибка при регистрации"
		sign["err"] = err.Error()

		at := make(map[string]any)
		at["user"] = user
		pg := page.NewPage(page.WithSignals(sign), page.WithAttrs(at))
		w.WriteHeader(http.StatusUnauthorized)
		page.Execute("index", "index", w, pg)
		return
	}
	http.SetCookie(w, ck)

	// Редирект на страницу верификации email

}
