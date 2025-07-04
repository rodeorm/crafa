package iteration

import (
	"context"
	"crafa/internal/core"
	"crafa/internal/http/page"
	"crafa/internal/logger"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

func UpdateGet(s SessionManager, i IterationStorager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := s.GetSession(r)
		if err != nil {
			logger.Log.Error("session",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		id, err := strconv.Atoi(r.URL.Query().Get("id"))

		// Редактировать итерацию может только администратор (уже проверяется в middle, на всякий случай и здесь)
		if err != nil || (session.User.Role.ID != core.RoleAdmin) {
			logger.Log.Error("id",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		if id == 0 {
			logger.Log.Error("id",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		iter := &core.Iteration{ID: id}
		at := make(map[string]any)
		ctx := context.TODO()

		err = i.SelectIteration(ctx, iter)
		if err != nil {
			logger.Log.Error("Iteration",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		at["iteration"] = iter

		pg := page.NewPage(page.WithAttrs(at), page.WithSession(session))
		page.Execute("iteration", "update", w, pg)
	}
}
