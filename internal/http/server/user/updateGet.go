package user

import (
	"context"
	"money/internal/core"
	"money/internal/http/page"
	"money/internal/logger"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

func UpdateGet(s SessionManager, u UserStorager, p ProjectStorager, t TeamStorager, rs RoleStorager) http.HandlerFunc {
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

		// Редактировать пользователя может либо сам пользователь, либо администратор
		if err != nil || (session.User.Role.ID != core.RoleAdmin && id != session.User.ID) {
			logger.Log.Error("id",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}
		user := &core.User{ID: id}
		at := make(map[string]any)
		ctx := context.TODO()

		err = u.SelectUser(ctx, user) // Получаем данные пользователя
		if err != nil {
			logger.Log.Error("User",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		// Получаем текущие проекты пользователя
		userProjects, err := p.SelectUserProjects(ctx, user)
		if err != nil {
			logger.Log.Error("userProjects",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}
		// Получаем текущие команды пользователя
		userTeams, err := t.SelectUserTeams(ctx, user)
		if err != nil {
			logger.Log.Error("userTeams",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		// Получаем возможные проекты для пользователя
		possibleProjects, err := p.SelectPossibleNewUserProjects(ctx, user)
		if err != nil {
			logger.Log.Error("possibleProjects",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		// Получаем возможные команды для пользователя
		possibleTeams, err := t.SelectPossibleNewUserTeams(ctx, user)
		if err != nil {
			logger.Log.Error("possibleTeams",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		err = rs.SelectRole(ctx, &user.Role)
		if err != nil {
			logger.Log.Error("Role",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		// Получаем возможные роли для пользователя
		possibleRoles, err := rs.SelectPossibleRoles(ctx)
		if err != nil {
			logger.Log.Error("possibleRoles",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		at["User"] = user
		at["PossibleRoles"] = possibleRoles
		at["PossibleProjects"] = possibleProjects
		at["PossibleTeams"] = possibleTeams
		at["UserProjects"] = userProjects
		at["UserTeams"] = userTeams

		pg := page.NewPage(page.WithAttrs(at), page.WithSession(session))
		page.Execute("user", "update", w, pg)

	}

}
