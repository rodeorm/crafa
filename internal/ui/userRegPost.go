package ui

import (
	"context"
	"log"
	"net/http"

	"money/internal/cookie"
	"money/internal/core"
)

func (h Handler) regPost(w http.ResponseWriter, r *http.Request) {

	user := core.User{
		Login:      r.FormValue("login"),
		Password:   r.FormValue("password"),
		Name:       r.FormValue("name"),
		Patronymic: r.FormValue("patronname"),
		Surname:    r.FormValue("familyname"),
		Email:      r.FormValue("email"),
		Phone:      r.FormValue("phonenumber"),
		Role:       core.Role{ID: 2}, // Пользователь создается с ролью "Неавторизованный пользователь"
	}

	ctx := context.TODO()
	var currentInformation sessionInformation
	err := h.Auth.RegUser(ctx, &user)
	if err != nil {
		log.Println("registerPost. ошибка в regPost на вставке нового пользователя", err)
		currentInformation.User = user
		currentInformation.Signal = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		ExecuteHTML("user", "reg", w, currentInformation)
		return
	}

	sessionID, err := h.Auth.AddSession(&user)
	if err != nil {
		log.Println("registerPost. ошибка в regPost на вставке в сессию", err)
		currentInformation.User = user
		currentInformation.Signal = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		ExecuteHTML("user", "reg", w, currentInformation)
		return
	}

	claims := cookie.CreateClaims(user.Login, sessionID, user.Role.ID, user.ID, user.Role.Name)
	http.SetCookie(w, cookie.PutTokenToCookie(claims))
	currentInformation.User = user
	ExecuteHTML("user", "confemail", w, currentInformation)
}
