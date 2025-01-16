package middle

import (
	"money/internal/crypt"
	"money/internal/http/cookie"
	"net/http"
)

func WithAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tkn, err := cookie.GetTokenFromRequest(r)
		if err != nil {
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}
		roleID, err := crypt.GetRoleIDFromTkn(tkn)
		if err != nil || roleID != 0 {
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}
		next.ServeHTTP(w, r)
	})
}
