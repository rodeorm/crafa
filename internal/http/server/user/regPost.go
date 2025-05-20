package user

import (
	"context"
	"money/internal/core"
	"money/internal/http/page"
	"net/http"
)

func RegPost(u UserStorager, c CookieManager, domain string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Получаем данные из формы
		user := core.User{
			Login:      r.FormValue("login"),
			Password:   r.FormValue("password"),
			Name:       r.FormValue("name"),
			PatronName: r.FormValue("patronname"),
			FamilyName: r.FormValue("familyname"),
			Email:      r.FormValue("email"),
			Phone:      r.FormValue("phonenumber"),
		}
		// Регистрируем пользователя. Получаем идентификатор пользователя и идентификатор сессии
		_, err := u.RegUser(context.TODO(), &user, domain)
		if err != nil {
			sign := make(map[string]string)
			sign["Russ"] = "Ошибка при регистрации"
			sign["Err"] = err.Error()
			at := make(map[string]any)
			at["User"] = user
			pg := page.NewPage(page.WithSignals(sign), page.WithAttrs(at))
			w.WriteHeader(http.StatusUnauthorized)
			page.Execute("user", "reg", w, pg)
			return
		}
		/*	Меняем логику. Пользователь регистрируется не сам, а через администратора. Код устаревает:
					// Создаем jwt-токен и сохраняем его в куках
					ck, err := c.NewCookieWithSession(session)
					if err != nil {
						sign := make(map[string]string)
						sign["russ"] = "Ошибка при регистрации"
						sign["err"] = err.Error()
						at := make(map[string]any)
						at["User"] = user
						pg := page.NewPage(page.WithSignals(sign), page.WithAttrs(at))
						w.WriteHeader(http.StatusUnauthorized)
						page.Execute("user", "reg", w, pg)
						return
					}

				http.SetCookie(w, ck)

			http.Redirect(w, r, "/user/wait", http.StatusTemporaryRedirect)
		*/

		http.Redirect(w, r, "/user/list", http.StatusSeeOther)
	}
}
