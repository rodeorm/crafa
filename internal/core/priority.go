package core

import "context"

type Priority struct {
	Level Level
	Name  string
	ID    int
}

type PriorityStorager interface {
	SelectPriority(context.Context, *Priority) error
	SelectAllPriorities(context.Context) ([]Priority, error)
	InsertPriority(ctx context.Context, a *Priority) error
	UpdatePriority(ctx context.Context, a *Priority) error
	SelectAllLevelPriorities(ctx context.Context, l *Level) error
}
