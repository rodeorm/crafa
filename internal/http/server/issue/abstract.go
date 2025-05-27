package issue

import (
	"context"
	"money/internal/core"
	"net/http"
)

type SessionManager interface {
	GetSession(r *http.Request) (*core.Session, error)
}

type IssueStorager interface {
	SelectIssue(context.Context, *core.Issue) error
	InsertIssue(context.Context, *core.Issue) error
	UpdateIssue(context.Context, *core.Issue) error
	DeleteIssue(context.Context, *core.Issue) error
}

type IssueSelecter interface {
	SelectAllIssues(context.Context) ([]core.Issue, error)
	SelectEpicIssues(context.Context, core.Epic) error
	SelectUserIssues(context.Context, core.User) error
	SelectUserStatusIssues(context.Context, core.User, core.Status) error
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
