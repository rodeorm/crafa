package area

import (
	"context"
	"crafa/internal/core"
	"crafa/internal/http/page"
	"crafa/internal/logger"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

func CreatePost(s SessionManager, a AreaStorager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := s.GetSession(r)
		if err != nil {
			logger.Log.Error("Session",
				zap.Error(err),
			)
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		levelID, err := strconv.Atoi(r.FormValue("levelid"))
		if err != nil {
			logger.Log.Error("AreaCreatePost. levelid",
				zap.Error(err))
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}
		Area := &core.Area{
			Name:  r.FormValue("name"),
			Level: core.Level{ID: levelID},
		}

		at := make(map[string]any)
		err = a.InsertArea(context.TODO(), Area)

		if err != nil {
			logger.Log.Error("InsertArea",
				zap.Error(err),
			)
			sign := make(map[string]string)
			sign["Russ"] = "Ошибка при создании категории"
			sign["Err"] = err.Error()
			pg := page.NewPage(page.WithSignals(sign), page.WithAttrs(at), page.WithSession(session))
			w.WriteHeader(http.StatusInternalServerError)
			page.Execute("area", "list", w, pg)
			return
		}

		http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
	}
}
