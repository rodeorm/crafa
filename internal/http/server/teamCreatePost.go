package server

import (
	"context"
	"money/internal/core"
	"money/internal/http/page"
	"money/internal/logger"
	"net/http"

	"go.uber.org/zap"
)

func (s *Server) teamCreatePost(w http.ResponseWriter, r *http.Request) {
	session, err := s.getSession(r)
	if err != nil {
		logger.Log.Error("Session",
			zap.Error(err),
		)
		http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
		return
	}

	Team := &core.Team{
		Name: r.FormValue("name"),
	}
	at := make(map[string]any)
	err = s.stgs.TeamStorager.InsertTeam(context.TODO(), Team)

	if err != nil {
		logger.Log.Error("InsertTeam",
			zap.Error(err),
		)
		sign := make(map[string]string)
		sign["Russ"] = "Ошибка при создании команды"
		sign["Err"] = err.Error()
		pg := page.NewPage(page.WithSignals(sign), page.WithAttrs(at), page.WithSession(session))
		page.Execute("team", "list", w, pg)
		return
	}
	http.Redirect(w, r, "/team/list", http.StatusSeeOther)
}
