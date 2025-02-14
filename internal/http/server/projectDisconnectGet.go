package server

import (
	"context"
	"fmt"
	"log"
	"money/internal/core"
	"money/internal/logger"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

func (s *Server) projectDisconnectGet(w http.ResponseWriter, r *http.Request) {
	session, err := s.getSession(r)
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
		log.Println("HERE")
		http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
	}
	ctx := context.TODO()

	err = s.stgs.DeleteUserProject(ctx, usr, pr)
	if err != nil {
		logger.Log.Error("DeleteUserProject",
			zap.Error(err),
		)
		http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/user/update?id=%d", userID), http.StatusSeeOther)

}
