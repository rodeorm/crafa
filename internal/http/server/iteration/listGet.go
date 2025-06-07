package iteration

import (
	"context"
	"crafa/internal/http/page"
	"crafa/internal/logger"
	"net/http"

	"go.uber.org/zap"
)

func ListGet(s SessionManager, i IterationStorager, l LevelStorager, t TeamStorager) http.HandlerFunc {
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

		iterations, err := i.SelectAllIterations(ctx)
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

		possibleTeams, err := t.SelectAllTeams(ctx)
		if err != nil {
			logger.Log.Error("possible teams",
				zap.Error(err))
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		at["PossibleLevels"] = possibleLevels
		at["PossibleTeams"] = possibleTeams
		at["Iterations"] = iterations
		pg := page.NewPage(page.WithSignals(sign), page.WithAttrs(at), page.WithSession(session))
		page.Execute("iteration", "list", w, pg)
	}

}
