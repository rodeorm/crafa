package server

import (
	"context"
	"money/internal/http/page"
	"money/internal/logger"
	"net/http"

	"go.uber.org/zap"
)

func (s *Server) categoryListGet(w http.ResponseWriter, r *http.Request) {
	session, err := s.getSession(r)
	if err != nil {
		logger.Log.Error("Session",
			zap.Error(err),
		)
		http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
		return
	}

	sign := make(map[string]string)
	at := make(map[string]any)

	ctx := context.TODO()

	categories, err := s.stgs.SelectAllCategories(ctx)
	if err != nil {
		logger.Log.Error("categories all",
			zap.Error(err))
		http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
		return
	}

	possibleLevels, err := s.stgs.SelectAllLevels(ctx)
	if err != nil {
		logger.Log.Error("possible levels",
			zap.Error(err))
		http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
		return
	}

	at["PossibleLevels"] = possibleLevels
	at["Categories"] = categories
	pg := page.NewPage(page.WithSignals(sign), page.WithAttrs(at), page.WithSession(session))
	page.Execute("category", "list", w, pg)
}
