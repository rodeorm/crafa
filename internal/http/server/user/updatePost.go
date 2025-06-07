package user

import (
	"context"
	"crafa/internal/core"
	"crafa/internal/http/page"
	"crafa/internal/logger"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

func UpdatePost(s SessionManager, u UserStorager, p ProjectStorager, t TeamStorager, rs RoleStorager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := s.GetSession(r)
		if err != nil {
			logger.Log.Error("Session",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		// Редактировать пользователя может либо сам пользователь, либо администратор
		if err != nil || (session.User.Role.ID != core.RoleAdmin && id != session.User.ID) {
			logger.Log.Error("id",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		roleID, err := strconv.Atoi(r.FormValue("roleid"))
		if err != nil {
			logger.Log.Error("role",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
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
		err = u.UpdateUser(context.TODO(), user)
		at["User"] = user

		if err != nil {
			logger.Log.Error("updateUser",
				zap.Error(err),
			)
			sign := make(map[string]string)
			sign["Russ"] = "Ошибка при обновлении"
			sign["Err"] = err.Error()
			pg := page.NewPage(page.WithSignals(sign), page.WithAttrs(at), page.WithSession(session))
			page.Execute("user", "update", w, pg)
			return
		}
		http.Redirect(w, r, "/user/list", http.StatusSeeOther)
	}
}
