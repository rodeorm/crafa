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
	InsertUserProject(ctx context.Context, userID, projectID int) error
	UpdateProject(context.Context, *core.Project) error
	SelectProject(context.Context, *core.Project) error
	SelectAllProjects(context.Context) ([]core.Project, error)
	SelectUserProjects(context.Context, *core.User) ([]core.Project, error)
	DeleteProject(context.Context, *core.Project) error
	DeleteUserProject(context.Context, *core.User, *core.Project) error
	SelectPossibleNewUserProjects(context.Context, *core.User) ([]core.Project, error)
	SelectAllProjectEpics(context.Context, *core.Project) ([]core.Epic, error)
}
