package core

import "context"

// Область
type Area struct {
	Level  Level
	ID     int
	Name   string
	Parent *Area
	Child  []Area
}

type AreaStorager interface {
	InsertArea(context.Context, *Area) error
	UpdateArea(context.Context, *Area) error
	SelectArea(context.Context, *Area) error
	SelectAllAreas(context.Context) ([]Area, error)
	SelectAllLevelAreas(context.Context, *Level) error
	DeleteArea(context.Context, *Area) error
}
