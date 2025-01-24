package server

import (
	"net/http"
)

func (s *Server) confirmGet(w http.ResponseWriter, r *http.Request) {
	// При подтверждении адреса электронной почты не обязательно быть авторизованным в системе
	// поэтому ошибка получения сессии не обрабатывается
	/*
		session, _ := s.getSession(r)

		values := r.URL.Query()

		email := values.Get("email")
		login := values.Get("login")
		otp := values.Get("otp")

			err := s.stgs.UserStorager.ConfirmUserEmail(context.TODO(), login, email, otp)
			if err != nil {
				http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			}

			at := make(map[string]any)
			at["login"] = login
			at["email"] = email
			page.Execute("user", "confirm", w, page.NewPage(page.WithSession(session), page.WithAttrs(at)))*/
}
