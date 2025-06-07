package team

import (
	"context"
	"crafa/internal/core"
	"crafa/internal/http/page"
	"crafa/internal/logger"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

func DeletePost(s SessionManager, t TeamStorager) http.HandlerFunc {
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
		// Удалять команду может только администратор
		if err != nil || session.User.Role.ID != core.RoleAdmin {
			logger.Log.Error("id",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		// Получаем данные из формы
		Team := &core.Team{
			ID: id,
		}
		at := make(map[string]any)
		err = t.DeleteTeam(context.TODO(), Team)
		at["Team"] = Team

		if err != nil {
			logger.Log.Error("updateTeam",
				zap.Error(err),
			)
			sign := make(map[string]string)
			sign["Russ"] = "Ошибка при обновлении"
			sign["Err"] = err.Error()
			pg := page.NewPage(page.WithSignals(sign), page.WithAttrs(at), page.WithSession(session))
			page.Execute("team", "update", w, pg)
			return
		}
		http.Redirect(w, r, "/team/list", http.StatusSeeOther)
	}

}
