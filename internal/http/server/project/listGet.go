package project

import (
	"context"
	"crafa/internal/core"
	"crafa/internal/http/page"
	"crafa/internal/logger"
	"net/http"

	"go.uber.org/zap"
)

func ListGet(s SessionManager, p ProjectSelecter) http.HandlerFunc {
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
		var projects []core.Project

		ctx := context.TODO()

		if session.User.Role.ID == core.RoleAdmin {
			projects, err = p.SelectAllProjects(ctx)
			if err != nil {
				logger.Log.Error("projects all",
					zap.Error(err))
				http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
				return
			}

		} else {
			projects, err = p.SelectUserProjects(ctx, &session.User)
			if err != nil {
				logger.Log.Error("projects user",
					zap.Error(err))
				http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
				return
			}

		}

		at["Projects"] = projects
		pg := page.NewPage(page.WithSignals(sign), page.WithAttrs(at), page.WithSession(session))
		switch session.User.Role.ID {
		case core.RoleAdmin:
			page.Execute("project", "adminList", w, pg)
		case core.RoleEmployee:
			page.Execute("project", "employeeList", w, pg)
		case core.RoleAuth:
			page.Execute("project", "authList", w, pg)
		}
	}

}
