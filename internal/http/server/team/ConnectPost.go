package team

import (
	"context"
	"crafa/internal/core"
	"crafa/internal/logger"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

func ConnectPost(s SessionManager, t TeamStorager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := s.GetSession(r)
		if err != nil {
			logger.Log.Error("Session",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}
		// Добавлять в проект может только администратор
		if session.User.Role.ID != core.RoleAdmin {
			logger.Log.Info("Role",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		userID, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			logger.Log.Error("id",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		TeamID, err := strconv.Atoi(r.FormValue("teamid"))
		if err != nil {
			logger.Log.Error("teamid",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		err = t.InsertUserTeams(context.TODO(), userID, TeamID)
		if err != nil {
			logger.Log.Error("InsertUserTeam",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}
		http.Redirect(w, r, r.Referer(), http.StatusSeeOther) // Редирект с сохранением метода StatusTemporaryRedirect
	}

}
