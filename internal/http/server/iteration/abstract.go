package iteration

import (
	"context"
	"money/internal/core"
	"net/http"
)

type SessionManager interface {
	GetSession(r *http.Request) (*core.Session, error)
}

type IterationStorager interface {
	InsertIteration(ctx context.Context, p *core.Iteration) error
	UpdateIteration(ctx context.Context, p *core.Iteration) error
	SelectIteration(ctx context.Context, p *core.Iteration) error
	SelectAllIterations(ctx context.Context) ([]core.Iteration, error)
	DeleteIteration(ctx context.Context, p *core.Iteration) error
	SelectPossibleLevelIterations(ctx context.Context, l *core.Level) ([]core.Iteration, error)
}

type LevelStorager interface {
	SelectLevel(context.Context, *core.Level) error
	SelectAllLevels(context.Context) ([]core.Level, error)
}

type TeamStorager interface {
	SelectAllTeams(ctx context.Context) ([]core.Team, error)
}
