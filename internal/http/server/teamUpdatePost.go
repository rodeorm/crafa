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

func (s *Server) teamUpdatePost(w http.ResponseWriter, r *http.Request) {
	session, err := s.getSession(r)
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
	prjct := &core.Team{
		ID:   id,
		Name: r.FormValue("name"),
	}
	at := make(map[string]any)
	err = s.stgs.TeamStorager.UpdateTeam(context.TODO(), prjct)
	at["Team"] = prjct

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
	http.Redirect(w, r, "/team/list", http.StatusTemporaryRedirect)
}
