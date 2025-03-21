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

func (s *Server) categoryCreatePost(w http.ResponseWriter, r *http.Request) {
	session, err := s.getSession(r)
	if err != nil {
		logger.Log.Error("Session",
			zap.Error(err),
		)
		http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
		return
	}

	levelID, err := strconv.Atoi(r.FormValue("levelid"))
	if err != nil {
		logger.Log.Error("categoryCreatePost. levelid",
			zap.Error(err))
		http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
		return
	}
	Category := &core.Category{
		Name:  r.FormValue("name"),
		Level: core.Level{ID: levelID},
	}

	at := make(map[string]any)
	err = s.stgs.CategoryStorager.InsertCategory(context.TODO(), Category)

	if err != nil {
		logger.Log.Error("InsertCategory",
			zap.Error(err),
		)
		sign := make(map[string]string)
		sign["Russ"] = "Ошибка при создании категории"
		sign["Err"] = err.Error()
		pg := page.NewPage(page.WithSignals(sign), page.WithAttrs(at), page.WithSession(session))
		page.Execute("category", "list", w, pg)
		return
	}

	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
}
