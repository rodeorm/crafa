package server

import (
	"context"
	"money/internal/core"
	"money/internal/http/page"
	"net/http"
	"strconv"
)

func (s *Server) userUpdatePost(w http.ResponseWriter, r *http.Request) {
	session, err := s.getSession(r)
	if err != nil {
		http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	// Редактировать пользователя может либо сам пользователь, либо администратор
	if err != nil || session.User.Role.ID != core.RoleAdmin || id != session.User.ID {
		http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
	}

	roleID, err := strconv.Atoi(r.FormValue("roleid"))
	if err != nil {
		http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
	}
	// Получаем данные из формы
	user := &core.User{
		ID:         id,
		Login:      r.FormValue("login"),
		Name:       r.FormValue("name"),
		PatronName: r.FormValue("patronname"),
		FamilyName: r.FormValue("familyname"),
		Email:      r.FormValue("email"),
		Phone:      r.FormValue("phonenumber"),
		Role:       core.Role{ID: roleID},
	}
	at := make(map[string]any)
	err = s.stgs.UserStorager.UpdateUser(context.TODO(), user)
	at["User"] = user

	if err != nil {
		sign := make(map[string]string)
		sign["Russ"] = "Ошибка при обновлении"
		sign["Err"] = err.Error()
		pg := page.NewPage(page.WithSignals(sign), page.WithAttrs(at), page.WithSession(session))
		page.Execute("user", "update", w, pg)
		return
	}

	pg := page.NewPage(page.WithAttrs(at), page.WithSession(session))
	page.Execute("user", "update", w, pg)

}
