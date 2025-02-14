package core

import "context"

// Область
type Area struct {
	Level
	Team
	ID int
}

type AreaStorager interface {
	AddArea(context.Context, *Area) error
	EditArea(context.Context, *Area) error
	SelectArea(context.Context, *Area) error
	SelectAllAreas(context.Context) ([]Area, error)
	DeleteArea(context.Context, *Area) error
}
