package user

import (
	"context"
	"money/internal/core"
	"money/internal/http/page"
	"net/http"
	"strconv"
)

func ConfirmGet(s SessionManager, u UserStorager, c CookieManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// При подтверждении адреса электронной почты не обязательно быть авторизованным в системе
		// поэтому ошибка получения сессии не обрабатывается
		session, _ := s.GetSession(r)
		values := r.URL.Query()

		id := values.Get("id")
		userID, err := strconv.Atoi(id)

		if err != nil {
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
		}
		otp := values.Get("otp")
		err = u.ConfirmUserEmail(context.TODO(), userID, otp)
		if err != nil {
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
		}

		at := make(map[string]any)
		at["id"] = userID

		if session != nil {
			session.Role.ID = core.RoleAuth
			// Создаем новый jwt-токен и сохраняем его в куках
			ck, _ := c.NewCookieWithSession(session)
			http.SetCookie(w, ck)
			page.Execute("user", "confirm", w, page.NewPage(page.WithAttrs(at), page.WithSession(session)))
			return
		}
		page.Execute("user", "confirm", w, page.NewPage(page.WithAttrs(at)))
	}
}
