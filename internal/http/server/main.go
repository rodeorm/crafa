package server

import (
	"log"
	"money/internal/core"
	"money/internal/http/page"
	"money/internal/logger"
	"net/http"

	"go.uber.org/zap"
)

func (s *Server) main(w http.ResponseWriter, r *http.Request) {

	session, err := s.getSession(r)
	if err != nil {
		log.Println("main", err)
		page.Execute("index", "index", w, page.NewPage())
		return
	}

	logger.Log.Info("Данные сессии",
		zap.Int("UserRoleID", session.User.Role.ID),
		zap.Int("SessionID", session.ID),
		zap.String("Login", session.User.Login),
	)

	switch session.User.Role.ID {
	case core.RoleAdmin:
		p := page.NewPage(page.WithSession(session))
		page.Execute("admin", "index", w, p)
	case core.RoleReg:
		http.Redirect(w, r, "/user/send", http.StatusTemporaryRedirect)
	case core.RoleAuth:
		p := page.NewPage(page.WithSession(session))
		page.Execute("main", "auth", w, p)
	}
}
