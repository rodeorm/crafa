package ui

import (
	"context"
	"log"
	"net/http"

	"money/internal/cookie"
	"money/internal/core"
)

func (h Handler) loginPost(w http.ResponseWriter, r *http.Request) {

	_, err := h.getSessionInformation(r)
	// Если пользователь не авторизован или его сессия уже истекла, то пробуем пройти аутентификацию
	if err != nil {

		user := core.User{Login: r.FormValue("login"), Password: r.FormValue("password")}
		ctx := context.TODO()
		success, err := h.Auth.AuthUser(ctx, &user)
		var currentInformation sessionInformation

		if !success {
			currentInformation.Signal = "Неправильное имя пользователя или пароль"
			w.WriteHeader(http.StatusUnauthorized)
			ExecuteHTML("index", "index", w, currentInformation)
			return
		}

		if err != nil {
			log.Println("loginPost", err)
			currentInformation.Signal = "Внутренняя ошибка сервера. Повторите попытку позднее"
			w.WriteHeader(http.StatusBadRequest)
			ExecuteHTML("index", "index", w, currentInformation)
			return
		}

		sessionID, err := h.Auth.AddSession(&user)
		if err != nil {
			log.Println("loginPost", err)
			currentInformation.Signal = "Внутренняя ошибка сервера. Повторите попытку позднее"
			w.WriteHeader(http.StatusBadRequest)
			ExecuteHTML("index", "index", w, currentInformation)
			return
		}

		claims := cookie.CreateClaims(user.Login, sessionID, user.Role.ID, user.ID, user.Role.Name)
		http.SetCookie(w, cookie.PutTokenToCookie(claims))
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

		return
	}

	// Если авторизован, то повторно не пробуем логиниться, сразу просто перенаправляемся на главную страницу
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
