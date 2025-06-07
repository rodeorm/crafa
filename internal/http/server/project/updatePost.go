package project

import (
	"context"
	"crafa/internal/core"
	"crafa/internal/http/page"
	"crafa/internal/logger"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

func UpdatePost(s SessionManager, p ProjectStorager) http.HandlerFunc {
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
		// Редактировать проект может только администратор
		if err != nil || (session.User.Role.ID != core.RoleAdmin) {
			logger.Log.Error("id",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		// Получаем данные из формы
		prjct := &core.Project{
			ID:   id,
			Name: r.FormValue("name"),
		}
		at := make(map[string]any)
		err = p.UpdateProject(context.TODO(), prjct)
		at["Project"] = prjct

		if err != nil {
			logger.Log.Error("updateProject",
				zap.Error(err),
			)
			sign := make(map[string]string)
			sign["Russ"] = "Ошибка при обновлении"
			sign["Err"] = err.Error()
			pg := page.NewPage(page.WithSignals(sign), page.WithAttrs(at), page.WithSession(session))
			page.Execute("project", "update", w, pg)
			return
		}
		http.Redirect(w, r, "/project/list", http.StatusTemporaryRedirect)
	}
}
