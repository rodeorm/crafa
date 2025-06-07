package priority

import (
	"context"
	"crafa/internal/core"
	"net/http"
)

type SessionManager interface {
	GetSession(r *http.Request) (*core.Session, error)
}

type LevelStorager interface {
	SelectLevel(context.Context, *core.Level) error
	SelectAllLevels(context.Context) ([]core.Level, error)
}

type TeamStorager interface {
	SelectAllTeams(ctx context.Context) ([]core.Team, error)
}

type PriorityStorager interface {
	SelectPriority(context.Context, *core.Priority) error
	SelectAllPriorities(context.Context) ([]core.Priority, error)
	InsertPriority(ctx context.Context, a *core.Priority) error
	UpdatePriority(ctx context.Context, a *core.Priority) error
	SelectAllLevelPriorities(ctx context.Context, l *core.Level) error
}
