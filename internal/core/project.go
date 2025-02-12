package core

import "context"

type Project struct {
	ID   int
	Name string

	Manager    User
	Supporter  User
	Maintainer User
}

type ProjectStorager interface {
	InsertProject(context.Context, *Project) error
	UpdateProject(context.Context, *Project) error
	SelectProject(context.Context, *Project) error
	SelectAllProjects(context.Context) ([]Project, error)
	SelectUserProjects(context.Context, User) ([]Project, error)
	DeleteProject(context.Context, *Project) error
}
