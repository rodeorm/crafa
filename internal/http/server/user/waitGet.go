package user

import (
	"context"
	"crafa/internal/http/page"
	"net/http"
)

func Wait(s SessionManager, u UserStorager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := s.GetSession(r)
		if err != nil {
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		u.SelectUser(context.TODO(), &session.User)
		page.Execute("user", "wait", w, page.NewPage(page.WithSession(session)))
	}
}
