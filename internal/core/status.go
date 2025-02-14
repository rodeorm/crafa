package core

import "context"

type Status struct {
	Level
	ID int
}

type StatusStorager interface {
	AddStatus(context.Context, *Status, *User) error
	EditStatus(context.Context, *Status, *User) error
	SelectStatus(context.Context, *Status, *User) error
	SelectAllStatuses(context.Context, *User) ([]Status, error)
	DeleteStatus(context.Context, *Status, *User) error
}
