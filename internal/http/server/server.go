package server

import (
	"crafa/internal/cfg"
	"crafa/internal/http/cookie"
	"crafa/internal/repo/cash"
	"crafa/internal/repo/postgres"
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
