package project

import (
	"context"
	"crafa/internal/core"
	"crafa/internal/logger"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

func DisconnectGet(s SessionManager, p ProjectStorager, u UserProjectManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := s.GetSession(r)
		if err != nil {
			logger.Log.Error("Session",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		userID, err := strconv.Atoi(r.URL.Query().Get("userid"))
		if err != nil {
			logger.Log.Error("userID",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}
		projectID, err := strconv.Atoi(r.URL.Query().Get("projectid"))
		if err != nil {
			logger.Log.Error("projectid",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}
		usr := &core.User{ID: userID}
		pr := &core.Project{ID: projectID}
		// Только сам пользователь или админ могут отвязать себя от проекта
		if userID != session.User.ID && session.Role.ID != core.RoleAdmin {
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
		}
		ctx := context.TODO()

		err = u.DeleteUserProject(ctx, usr, pr)
		if err != nil {
			logger.Log.Error("DeleteUserProject",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
	}
}
