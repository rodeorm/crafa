package status

import (
	"context"
	"money/internal/http/page"
	"money/internal/logger"
	"net/http"

	"go.uber.org/zap"
)

func ListGet(s SessionManager, st StatusStorager, l LevelStorager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := s.GetSession(r)
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

		Statuses, err := st.SelectAllStatuses(ctx)
		if err != nil {
			logger.Log.Error("categories all",
				zap.Error(err))
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		possibleLevels, err := l.SelectAllLevels(ctx)
		if err != nil {
			logger.Log.Error("possible levels",
				zap.Error(err))
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		at["PossibleLevels"] = possibleLevels
		at["Statuses"] = Statuses
		pg := page.NewPage(page.WithSignals(sign), page.WithAttrs(at), page.WithSession(session))
		page.Execute("status", "list", w, pg)
	}
}
