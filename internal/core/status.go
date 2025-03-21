package core

import (
	"context"
)

type Status struct {
	Level Level
	Name  string
	ID    int
}

type StatusStorager interface {
	SelectStatus(context.Context, *Status) error
	SelectAllStatuses(context.Context) ([]Status, error)
	InsertStatus(ctx context.Context, a *Status) error
	UpdateStatus(ctx context.Context, a *Status) error
	SelectAllLevelStatuses(ctx context.Context, l *Level) error
}
