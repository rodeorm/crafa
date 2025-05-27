package project

import (
	"context"
	"money/internal/core"
	"money/internal/http/page"
	"money/internal/logger"
	"net/http"
	"strconv"
)

func ViewGet(s SessionManager, p ProjectStorager, u ProjectUserSelecter, e ProjectEpicSelecter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := s.GetSession(r)
		if err != nil {
			logger.Sugar.Error(err)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		id, err := strconv.Atoi(r.URL.Query().Get("id"))

		// Редактировать проект может только администратор (уже проверяется в middle, на всякий случай и здесь)
		if err != nil || (session.User.Role.ID != core.RoleAdmin) {
			logger.Sugar.Error(err)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}
		prjct := &core.Project{ID: id}
		at := make(map[string]any)
		ctx := context.TODO()

		err = p.SelectProject(ctx, prjct)
		if err != nil {
			logger.Sugar.Error(err)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		err = e.SelectProjectStatusEpics(ctx, prjct, nil)
		if err != nil {
			logger.Sugar.Error(err)
			//http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		err = u.SelectProjectUsers(ctx, prjct)
		if err != nil {
			logger.Sugar.Error(err)
			//http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}
		at["Project"] = prjct

		pg := page.NewPage(page.WithAttrs(at), page.WithSession(session))
		err = page.Execute("project", "update", w, pg)
		if err != nil {
			logger.Sugar.Error(err)
			//http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}
	}
}
