package user

import (
	"context"
	"crafa/internal/core"
	"crafa/internal/http/page"
	"crafa/internal/logger"
	"net/http"
	"strconv"
	"time"
)

func VerifyPost(s SessionManager, u UserStorager, c CookieManager, liveTime time.Duration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		otp := r.FormValue("otp")
		usr := &core.User{ID: id}
		at := make(map[string]any)
		at["User"] = usr

		session, err := u.AdvAuthUser(context.TODO(), usr, otp, liveTime)
		if err != nil {
			logger.Sugar.Error(err)
			sign := make(map[string]string)
			sign["russ"] = "Неправильный код подтверждения"
			sign["err"] = err.Error()
			pg := page.NewPage(page.WithSignals(sign), page.WithAttrs(at))
			w.WriteHeader(http.StatusUnauthorized)
			page.Execute("user", "verify", w, pg)
			return
		}

		ck, err := c.NewCookieWithSession(session)
		if err != nil {
			sign := make(map[string]string)
			sign["russ"] = "Ошибка при аутентификации"
			sign["err"] = err.Error()
			pg := page.NewPage(page.WithSignals(sign), page.WithAttrs(at))
			w.WriteHeader(http.StatusUnauthorized)
			page.Execute("user", "verify", w, pg)
			return
		}

		http.SetCookie(w, ck)
		http.Redirect(w, r, "/main", http.StatusTemporaryRedirect)
	}

}
