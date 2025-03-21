package server

import (
	"context"
	"money/internal/core"
	"money/internal/logger"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

func (s *Server) teamDisconnectGet(w http.ResponseWriter, r *http.Request) {
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
	TeamID, err := strconv.Atoi(r.URL.Query().Get("teamid"))
	if err != nil {
		logger.Log.Error("teamid",
			zap.Error(err),
		)
		http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
		return
	}
	usr := &core.User{ID: userID}
	pr := &core.Team{ID: TeamID}
	// Только сам пользователь или админ могут отвязать себя от команды
	if userID != session.User.ID && session.Role.ID != core.RoleAdmin {
		http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
	}
	ctx := context.TODO()

	err = s.stgs.DeleteUserTeam(ctx, usr, pr)
	if err != nil {
		logger.Log.Error("DeleteUserTeam",
			zap.Error(err),
		)
		http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
		return
	}

	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)

}
