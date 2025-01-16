package server

import (
	"money/internal/cfg"
	"money/internal/core"
	"net/http"
)

type Server struct {
	srv      *http.Server
	storages *core.Storage
	cfg      *cfg.ServerConfig

	exit chan struct{}
}
