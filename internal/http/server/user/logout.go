package user

import (
	"net/http"
)

func LogOut(w http.ResponseWriter, r *http.Request) {
	ck := &http.Cookie{
		Name:  "token",
		Value: "",
		Path:  "/",
	}
	http.SetCookie(w, ck)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
