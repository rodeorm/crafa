package middle

import (
	"crafa/internal/core"
	"crafa/internal/crypt"
	"crafa/internal/http/cookie"
	"net/http"
)

func WithAdmin(jwtKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tkn, err := cookie.GetTokenFromRequest(r)
			if err != nil {
				http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
				return
			}
			roleID, err := crypt.GetRoleIDFromTkn(tkn, jwtKey)
			if err != nil || roleID != core.RoleAdmin {
				http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
