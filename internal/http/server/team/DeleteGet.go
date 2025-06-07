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

func DeleteGet(s SessionManager, t TeamStorager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := s.GetSession(r)

		if err != nil {
			logger.Log.Error("session",
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

		team := &core.Team{ID: id}
		at := make(map[string]any)
		ctx := context.TODO()

		err = t.SelectTeam(ctx, team) // Получаем данные команды
		if err != nil {
			logger.Log.Error("SelectTeam",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		at["Team"] = team

		pg := page.NewPage(page.WithAttrs(at), page.WithSession(session))
		page.Execute("team", "delete", w, pg)

	}
}
