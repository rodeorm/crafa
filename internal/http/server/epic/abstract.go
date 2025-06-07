package epic

import (
	"context"
	"crafa/internal/core"
	"net/http"
)

type SessionManager interface {
	GetSession(r *http.Request) (*core.Session, error)
}

type EpicStorager interface {
	SelectEpic(context.Context, *core.Epic) error
	InsertEpic(context.Context, *core.Epic) error
	UpdateEpic(context.Context, *core.Epic) error
	DeleteEpic(context.Context, *core.Epic) error
}

type EpicSelecter interface {
	SelectAllEpics(context.Context) ([]core.Epic, error)
	SelectFilteredEpics(context.Context, []core.Project, []core.Status, []core.User) ([]core.Epic, error)
}
