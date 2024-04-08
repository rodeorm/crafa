package ui

import (
	"context"
	"log"
	"money/internal/core"
	"net/http"
)

func (h Handler) index(w http.ResponseWriter, r *http.Request) {
	currentInformation, err := h.getCurrentUserInformation(r)

	if err != nil {
		log.Println("index", err)
		ExecuteHTML("index", "index", w, *currentInformation)
		return
	}
	ctx := context.TODO()

	switch currentInformation.User.Role.ID {
	case Admin:
		currentInformation.Attribute = core.Book{}
		ExecuteHTML("admin", "accountList", w, *currentInformation)
	case UnauthorisedUser:
		h.Auth.SelectUser(ctx, &currentInformation.User)
		ExecuteHTML("index", "indexAuth", w, *currentInformation)
	case AuthorisedUser:
		currentInformation.Attribute = core.Book{}
		ExecuteHTML("index", "indexAuth", w, *currentInformation)
	}

}
