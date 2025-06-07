package user

import (
	"context"
	"net/http"

	"crafa/internal/core"
	"crafa/internal/http/page"
)

func LoginPost(u UserStorager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &core.User{Login: r.FormValue("login"), Password: r.FormValue("password")}
		ctx := context.Background()

		err := u.BaseAuthUser(ctx, user)
		if err != nil {
			sign := make(map[string]string)
			sign["Russ"] = "неправильное имя пользователя или пароль"
			sign["Err"] = err.Error()
			pg := page.NewPage(page.WithSignals(sign))
			w.WriteHeader(http.StatusUnauthorized)
			page.Execute("index", "index", w, pg)
			return
		}

		at := make(map[string]any)
		at["User"] = user

		page.Execute("user", "verify", w, page.NewPage(page.WithAttrs(at)))
	}

}
