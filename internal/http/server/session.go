package server

import (
	"net/http"

	"crafa/internal/core"
	"crafa/internal/http/cookie"
)

func (s *Server) GetSession(r *http.Request) (*core.Session, error) {
	tkn, err := cookie.GetTokenFromRequest(r)
	if err != nil {
		//log.Println("getSession 1", err)
		return nil, err
	}

	sn, err := core.GetSessionFromTkn(tkn, s.cfg.JWTKey)
	if err != nil {
		//log.Println("getSession 2", err)
		return nil, err
	}
	return sn, nil
}
