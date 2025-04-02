package server

import (
	"money/internal/cfg"
	"money/internal/core"
	"money/internal/http/cookie"
	"money/internal/repo/postgres"
	"net/http"
)

type Server struct {
	srv  *http.Server
	stgs *core.Storage
	cfg  *cfg.Config
	ps   *postgres.PostgresStorage
	cm   *cookie.CookieManager

	exit chan struct{}
}
