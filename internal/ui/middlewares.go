package ui

import (
	"net/http"

	"money/internal/cookie"
)

// AuthMiddleware выполняется для проверки аутентифицирован ли пользователь
func (h Handler) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := cookie.GetClaimsFromCookie(r)
		if err != nil {
			// log.Println(r.RequestURI, r.Method, "Забраковали на проверке в auth на шаге: проверка сохраненных куков")
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}
		_, err = h.Auth.SelectActiveSession(claims.UserID, claims.SessionID)
		if err != nil {
			// log.Println("Забраковали на проверке в auth на шаге: проверка активных сессий: ", claims.Login, claims.RoleID)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// AdminMiddleware выполняется для проверки имеет ли пользователь роль администратора
func (h Handler) adminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := cookie.GetClaimsFromCookie(r)
		if err != nil {
			//	log.Println("Забраковали на проверке в admin на шаге: проверка сохраненных куков")
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		if claims.RoleID != Admin {
			//	log.Println(r.RequestURI, r.Method, "Забраковали на проверке в admin на шаге: проверка роли \"Администратор\": ", claims.Login, claims.RoleID)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// LogMiddleware выполняется для логирования действий пользователя
func (h Handler) logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log.Print(r.RequestURI, r.Method)
		next.ServeHTTP(w, r)
	})
}
