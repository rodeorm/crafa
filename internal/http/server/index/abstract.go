package index

import (
	"money/internal/core"
	"net/http"
)

type SessionManager interface {
	GetSession(r *http.Request) (*core.Session, error)
}
