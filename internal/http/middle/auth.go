package middle

import (
	"money/internal/core"
	"money/internal/crypt"
	"money/internal/http/cookie"
	"net/http"
)

func WithAuth(jwtKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tkn, err := cookie.GetTokenFromRequest(r)
			if err != nil {
				http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
				return
			}
			roleID, err := crypt.GetRoleIDFromTkn(tkn, jwtKey)
			if err != nil || (roleID != core.RoleAuth && roleID != core.RoleAdmin) {
				http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
