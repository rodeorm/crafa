package core

import "context"

type Project struct {
	ID   int
	Name string

	IssueQty     int
	OpenIssueQty int

	Manager    User
	Supporter  User
	Maintainer User
}

type ProjectStorager interface {
	InsertProject(context.Context, *Project) error
	InsertUserProject(ctx context.Context, userID, projectID int) error
	UpdateProject(context.Context, *Project) error
	SelectProject(context.Context, *Project) error
	SelectAllProjects(context.Context) ([]Project, error)
	SelectUserProjects(context.Context, *User) ([]Project, error)
	DeleteProject(context.Context, *Project) error
	DeleteUserProject(context.Context, *User, *Project) error
	SelectPossibleNewUserProjects(context.Context, *User) ([]Project, error)
}
