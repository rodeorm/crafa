package middle

import (
	"net/http"
)

func WithLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log.Print(r.RequestURI, r.Method)
		next.ServeHTTP(w, r)
	})
}
