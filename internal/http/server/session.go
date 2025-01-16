package server

import (
	"net/http"

	"money/internal/core"
	"money/internal/crypt"
	"money/internal/http/cookie"
)

func (s *Server) getSession(r *http.Request) (*core.Session, error) {
	tkn, err := cookie.GetTokenFromRequest(r)
	if err != nil {
		return nil, err
	}

	sn, err := crypt.GetSessionFromTkn(tkn)
	if err != nil {
		return nil, err
	}
	return sn, nil
}
