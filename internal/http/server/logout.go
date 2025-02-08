package server

import (
	"money/internal/http/cookie"
	"money/internal/http/page"
	"money/internal/logger"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func (s *Server) logOut(w http.ResponseWriter, r *http.Request) {
	session, err := s.getSession(r)

	logger.Log.Info("Данные сессии",
		zap.Int("UserRoleID", session.User.Role.ID),
		zap.Int("SessionID", session.ID),
		zap.String("Login", session.User.Login),
	)

	p := page.NewPage(page.WithSession(session))

	if err != nil {
		logger.Log.Error("Данные сессии",
			zap.String("err", err.Error()),
		)
		page.Execute("index", "index", w, p)
		return
	}

	session.LogoutTime.Time = time.Now()
	session.LogoutTime.Valid = true

	/* err = s.stgs.SessionStorager.UpdateSession(context.TODO(), session)
	if err != nil {
		logger.Log.Error("UpdateSession",
			zap.String("err", err.Error()),
		)
		page.Execute("index", "index", w, p)
		return
	}
	*/
	http.SetCookie(w, cookie.RemoveTokenFromCookie())
	page.Execute("index", "index", w, page.NewPage())
}
