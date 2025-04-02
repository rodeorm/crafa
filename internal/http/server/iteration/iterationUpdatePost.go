package iteration

import (
	"context"
	"database/sql"
	"money/internal/core"
	"money/internal/http/page"
	"money/internal/logger"
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

func iterationUpdatePost(s SessionManager, i IterationStorager) http.HandlerFunc {
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
			logger.Log.Error("iterationUpdatePost. levelid",
				zap.Error(err))
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		parentID, err := strconv.Atoi(r.FormValue("parentid"))
		if err != nil {
			logger.Log.Error("iterationUpdatePost. parentid",
				zap.Error(err))
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}

		year, err := strconv.Atoi(r.FormValue("year"))
		if err != nil {
			logger.Log.Error("iterationUpdatePost. year",
				zap.Error(err))
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}
		month, err := strconv.Atoi(r.FormValue("month"))
		if err != nil {
			logger.Log.Error("IterationUpdatePost. month",
				zap.Error(err))
			http.Redirect(w, r, "/forbidden", http.StatusTemporaryRedirect)
			return
		}
		monthNullable := sql.NullInt32{}
		if month != 0 {
			monthNullable.Int32 = int32(month)
			monthNullable.Valid = true
		}

		iter := &core.Iteration{
			Name:   r.FormValue("name"),
			Level:  core.Level{ID: levelID},
			Parent: &core.Iteration{ID: parentID},
			Year:   year,
			Month:  monthNullable,
		}

		if year < 2000 && year > 2050 {
			logger.Log.Error("InsertIteration",
				zap.Error(err),
			)
			sign := make(map[string]string)
			sign["Russ"] = "Ошибка при создании итерации"
			sign["Err"] = "Год должен быть между 2000 до 2050"
			at := make(map[string]any)
			at["iter"] = iter
			pg := page.NewPage(page.WithSignals(sign), page.WithAttrs(at), page.WithSession(session))
			page.Execute("iteration", "update", w, pg)
			return

		}

		err = i.UpdateIteration(context.TODO(), iter)

		if err != nil {
			logger.Log.Error("InsertIteration",
				zap.Error(err),
			)
			sign := make(map[string]string)
			sign["Russ"] = "Ошибка при создании категории"
			sign["Err"] = err.Error()
			at := make(map[string]any)
			at["iter"] = iter

			pg := page.NewPage(page.WithSignals(sign), page.WithAttrs(at), page.WithSession(session))
			page.Execute("iteration", "update", w, pg)
			return
		}

		http.Redirect(w, r, "/iteration/list", http.StatusSeeOther)
	}
}
