package index

import (
	"crafa/internal/core"
	"crafa/internal/http/page"
	"net/http"
)

func MainMenu(s SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := s.GetSession(r)
		if err != nil {
			page.Execute("index", "index", w, page.NewPage())
			return
		}

		switch session.User.Role.ID {
		case core.RoleAdmin:
			p := page.NewPage(page.WithSession(session))
			page.Execute("admin", "main", w, p)
		case core.RoleReg:
			http.Redirect(w, r, "/user/send", http.StatusTemporaryRedirect)
		case core.RoleAuth:
			p := page.NewPage(page.WithSession(session))
			page.Execute("auth", "main", w, p)
		case core.RoleEmployee:
			p := page.NewPage(page.WithSession(session))
			page.Execute("employee", "main", w, p)
		}
	}
}
