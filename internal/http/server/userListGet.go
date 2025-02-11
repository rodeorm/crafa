package server

import (
	"context"
	"money/internal/core"
	"money/internal/http/page"
	"net/http"
)

func (s *Server) userListGet(w http.ResponseWriter, r *http.Request) {
	session, err := s.getSession(r)
	if err != nil || session.User.Role.ID != core.RoleAdmin {
		http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
	}
	sign := make(map[string]string)
	at := make(map[string]any)
	users, err := s.stgs.UserStorager.SelectAllUsers(context.TODO())
	if err != nil {
		sign["Russ"] = "внутренняя ошибка"
		sign["Err"] = err.Error()
	}
	at["Users"] = users
	pg := page.NewPage(page.WithSignals(sign), page.WithAttrs(at), page.WithSession(session))
	page.Execute("user", "list", w, pg)
}
