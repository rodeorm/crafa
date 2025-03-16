package server

import (
	"context"
	"money/internal/core"
	"money/internal/http/page"
	"money/internal/logger"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

func (s *Server) teamUpdateGet(w http.ResponseWriter, r *http.Request) {
	session, err := s.getSession(r)
	if err != nil {
		logger.Log.Error("session",
			zap.Error(err),
		)
		http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	// Редактировать команду может только администратор (уже проверяется в middle, на всякий случай и здесь)
	if err != nil || (session.User.Role.ID != core.RoleAdmin) {
		logger.Log.Error("id",
			zap.Error(err),
		)
		http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
		return
	}
	prjct := &core.Team{ID: id}
	at := make(map[string]any)
	ctx := context.TODO()

	err = s.stgs.TeamStorager.SelectTeam(ctx, prjct)
	if err != nil {
		logger.Log.Error("Team",
			zap.Error(err),
		)
		http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
		return
	}

	at["Team"] = prjct

	pg := page.NewPage(page.WithAttrs(at), page.WithSession(session))
	page.Execute("user", "update", w, pg)

}
