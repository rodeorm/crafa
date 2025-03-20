package core

import (
	"context"
	"database/sql"
)

type Iteration struct {
	Level Level
	Name  string

	ID    int
	Year  int
	Month sql.NullInt32

	Parent *Iteration
	Child  []Iteration
}

type IterationStorager interface {
	InsertIteration(ctx context.Context, p *Iteration) error
	UpdateIteration(ctx context.Context, p *Iteration) error
	SelectIteration(ctx context.Context, p *Iteration) error
	SelectAllIterations(ctx context.Context) ([]Iteration, error)
	DeleteIteration(ctx context.Context, p *Iteration) error
	SelectPossibleLevelIterations(ctx context.Context, l *Level) ([]Iteration, error)
}
