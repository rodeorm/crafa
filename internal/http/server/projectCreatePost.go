package server

import (
	"context"
	"money/internal/core"
	"money/internal/http/page"
	"money/internal/logger"
	"net/http"

	"go.uber.org/zap"
)

func (s *Server) projectCreatePost(w http.ResponseWriter, r *http.Request) {
	session, err := s.getSession(r)
	if err != nil {
		logger.Log.Error("Session",
			zap.Error(err),
		)
		http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
		return
	}

	project := &core.Project{
		Name: r.FormValue("name"),
	}
	at := make(map[string]any)
	err = s.stgs.ProjectStorager.InsertProject(context.TODO(), project)

	if err != nil {
		logger.Log.Error("InsertProject",
			zap.Error(err),
		)
		sign := make(map[string]string)
		sign["Russ"] = "Ошибка при создании проекта"
		sign["Err"] = err.Error()
		pg := page.NewPage(page.WithSignals(sign), page.WithAttrs(at), page.WithSession(session))
		page.Execute("project", "list", w, pg)
		return
	}
	http.Redirect(w, r, "/project/list", http.StatusSeeOther)
}
