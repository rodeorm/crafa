package server

import (
	"context"
	"net/http"

	"money/internal/core"
	"money/internal/http/page"
)

func (s *Server) regPost(w http.ResponseWriter, r *http.Request) {

	user := core.User{
		Login:      r.FormValue("login"),
		Password:   r.FormValue("password"),
		Name:       r.FormValue("name"),
		Patronymic: r.FormValue("patronname"),
		Surname:    r.FormValue("familyname"),
		Email:      r.FormValue("email"),
		Phone:      r.FormValue("phonenumber"),
		Role:       core.Role{ID: core.Auth},
	}

	ctx := context.TODO()

	err := s.storages.RegUser(ctx, &user)
	if err != nil {

	}

	session, err := s.storages.AddSession(ctx, &user)
	if err != nil {

	}

	s.storages.AddMessage(ctx, &core.Message{Login: session.Login, Destination: session.Email})
	page.Execute("email", "verify", w, session, nil)
}
