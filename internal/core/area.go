package core

import "context"

type Area struct {
	Level
	Team
	ID int
}

type AreaStorager interface {
	AddArea(context.Context, *Area, *User) error
	EditArea(context.Context, *Area, *User) error
	SelectArea(context.Context, *Area, *User) error
	SelectAllAreas(context.Context, *User) ([]Area, error)
	DeleteArea(context.Context, *Area, *User) error
}
