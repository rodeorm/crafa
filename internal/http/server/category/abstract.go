package category

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

type CategoryStorager interface {
	InsertCategory(ctx context.Context, c *core.Category) error
	UpdateCategory(ctx context.Context, c *core.Category) error
	SelectCategory(ctx context.Context, c *core.Category) error
	SelectAllCategories(ctx context.Context) ([]core.Category, error)
	SelectAllLevelCategories(ctx context.Context, l *core.Level) error
	DeleteCategory(ctx context.Context, c *core.Category) error
}
