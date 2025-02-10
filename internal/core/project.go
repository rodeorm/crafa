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
	AddProject(context.Context, *Project, *User) error
	EditProject(context.Context, *Project, *User) error
	SelectProject(context.Context, *Project, *User) error
	SelectAllProjects(context.Context, *User) ([]Project, error)
	DeleteProject(context.Context, *Project, *User) error
}
