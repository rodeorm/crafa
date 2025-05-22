package project

import (
	"context"
	"money/internal/core"
	"net/http"
)

type SessionManager interface {
	GetSession(r *http.Request) (*core.Session, error)
}

type ProjectStorager interface { //TODO: разбить на два интерфейса
	InsertProject(context.Context, *core.Project) error
	UpdateProject(context.Context, *core.Project) error
	SelectProject(context.Context, *core.Project) error
	DeleteProject(context.Context, *core.Project) error
}

type ProjectSelecter interface {
	SelectAllProjects(context.Context) ([]core.Project, error)
	//SelectStatusProjects(context.Context, core.Status) ([]core.Status, error)
	SelectUserProjects(context.Context, *core.User) ([]core.Project, error)
}

type ProjectEpicSelecter interface {
	SelectProjectStatusEpics(context.Context, *core.Project, *core.Status) error
}

type ProjectUserSelecter interface {
	SelectProjectUsers(context.Context, *core.Project) error
}

type UserProjectManager interface {
	InsertUserProject(ctx context.Context, userID, projectID int) error
	SelectUserProject(context.Context, *core.Project, *core.User) error
	DeleteUserProject(context.Context, *core.User, *core.Project) error
	SelectPossibleNewUserProjects(context.Context, *core.User) ([]core.Project, error)
}
