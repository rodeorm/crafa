package team

import (
	"context"
	"money/internal/core"
	"money/internal/http/page"
	"money/internal/logger"
	"net/http"

	"go.uber.org/zap"
)

func ListGet(s SessionManager, t TeamStorager) http.HandlerFunc {
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
		var Teams []core.Team

		ctx := context.TODO()

		if session.User.Role.ID == core.RoleAdmin {
			Teams, err = t.SelectAllTeams(ctx)
			if err != nil {
				logger.Log.Error("Teams all",
					zap.Error(err))
				http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
				return
			}

		} else {
			Teams, err = t.SelectUserTeams(ctx, &session.User)
			if err != nil {
				logger.Log.Error("Teams user",
					zap.Error(err))
				http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
				return
			}

		}

		at["Teams"] = Teams
		pg := page.NewPage(page.WithSignals(sign), page.WithAttrs(at), page.WithSession(session))
		switch session.User.Role.ID {
		case core.RoleAdmin:
			page.Execute("team", "adminList", w, pg)
		case core.RoleEmployee:
			page.Execute("team", "employeeList", w, pg)
		case core.RoleAuth:
			page.Execute("team", "authList", w, pg)
		}
	}
}
