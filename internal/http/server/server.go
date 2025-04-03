package server

import (
	"money/internal/cash"
	"money/internal/cfg"
	"money/internal/http/cookie"
	"money/internal/repo/postgres"
	"net/http"
)

type Server struct {
	srv  *http.Server
	cfg  *cfg.Config
	ps   *postgres.PostgresStorage
	cm   *cookie.CookieManager
	cash *cash.CashStorage
	exit chan struct{}
}
