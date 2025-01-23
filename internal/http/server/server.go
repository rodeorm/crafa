package server

import (
	"money/internal/cfg"
	"money/internal/core"
	"net/http"
)

type Server struct {
	srv  *http.Server
	stgs *core.Storage
	cfg  *cfg.Config

	exit chan struct{}
}
