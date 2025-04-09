package project

import (
	"context"
	"money/internal/core"
	"money/internal/http/page"
	"money/internal/logger"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

func ViewGet(s SessionManager, p ProjectStorager, u UserProjectManager) http.HandlerFunc {
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
		if err != nil {
			logger.Log.Error("Project",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}
		prjct := &core.Project{ID: id}
		at := make(map[string]any)
		ctx := context.TODO()

		err = u.SelectUserProject(ctx, prjct, &session.User)
		if err != nil {
			logger.Log.Error("Project",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		at["Project"] = prjct
		pg := page.NewPage(page.WithAttrs(at), page.WithSession(session))
		page.Execute("project", "view", w, pg)
	}
}
