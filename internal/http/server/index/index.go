package index

import (
	"money/internal/core"
	"money/internal/http/page"
	"net/http"
)

func Index(s SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := s.GetSession(r)
		if err != nil {
			page.Execute("index", "index", w, page.NewPage())
			return
		}

		p := page.NewPage(page.WithSession(session))
		switch session.User.Role.ID {
		case core.RoleGuest:
			page.Execute("index", "index", w, p)
		case core.RoleReg:
			http.Redirect(w, r, "/user/send", http.StatusTemporaryRedirect)
		case core.RoleAdmin:
			http.Redirect(w, r, "/main", http.StatusTemporaryRedirect)
		case core.RoleAuth:
			http.Redirect(w, r, "/main", http.StatusTemporaryRedirect)
		case core.RoleEmployee:
			http.Redirect(w, r, "/main", http.StatusTemporaryRedirect)
		}
	}
}
