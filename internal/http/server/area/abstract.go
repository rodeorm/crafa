package area

import (
	"context"
	"money/internal/core"
	"net/http"
)

type SessionManager interface {
	GetSession(r *http.Request) (*core.Session, error)
}

type AreaStorager interface {
	InsertArea(context.Context, *core.Area) error
	UpdateArea(context.Context, *core.Area) error
	SelectArea(context.Context, *core.Area) error
	SelectAllAreas(context.Context) ([]core.Area, error)
	SelectAllLevelAreas(context.Context, *core.Level) error
	DeleteArea(context.Context, *core.Area) error
}
