package core

import "context"

type Iteration struct {
	Level
	Name   string
	ID     int
	Year   int
	Month  int
	Parent *Iteration
	Child  *Iteration
}

type IterationStorager interface {
	AddIteration(context.Context, *Iteration, *User) error
	EditIteration(context.Context, *Iteration, *User) error
	SelectIteration(context.Context, *Iteration, *User) error
	SelectAllIterations(context.Context, *User) ([]Iteration, error)
	DeleteIteration(context.Context, *Iteration, *User) error
}
