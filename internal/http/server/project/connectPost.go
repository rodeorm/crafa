package project

import (
	"context"
	"money/internal/core"
	"money/internal/logger"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

func ConnectPost(s SessionManager, p ProjectStorager) http.HandlerFunc {
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

		projectID, err := strconv.Atoi(r.FormValue("projectid"))
		if err != nil {
			logger.Log.Error("projectid",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		err = p.InsertUserProject(context.TODO(), userID, projectID)
		if err != nil {
			logger.Log.Error("InsertUserProject",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}
		http.Redirect(w, r, r.Referer(), http.StatusSeeOther) // Редирект с сохранением метода StatusTemporaryRedirect
	}
}
