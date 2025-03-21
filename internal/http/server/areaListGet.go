package server

import (
	"context"
	"money/internal/http/page"
	"money/internal/logger"
	"net/http"

	"go.uber.org/zap"
)

func (s *Server) areaListGet(w http.ResponseWriter, r *http.Request) {
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

	areas, err := s.stgs.SelectAllAreas(ctx)
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
	at["Areas"] = areas
	pg := page.NewPage(page.WithSignals(sign), page.WithAttrs(at), page.WithSession(session))
	page.Execute("area", "list", w, pg)
}
