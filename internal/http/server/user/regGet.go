package user

import (
	"crafa/internal/http/page"
	"net/http"
)

func RegGet(w http.ResponseWriter, r *http.Request) {
	page.Execute("user", "reg", w, page.NewPage())
}
