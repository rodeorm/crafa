package index

import (
	"money/internal/http/page"
	"net/http"
)

func Forbidden(s SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := s.GetSession(r)
		page.Execute("index", "forbidden", w, page.NewPage(page.WithSession(session)))
	}
}
