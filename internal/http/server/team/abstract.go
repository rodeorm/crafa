package team

import (
	"context"
	"crafa/internal/core"
	"net/http"
)

type SessionManager interface {
	GetSession(r *http.Request) (*core.Session, error)
}

type TeamStorager interface {
	InsertTeam(ctx context.Context, p *core.Team) error
	SelectTeam(ctx context.Context, p *core.Team) error
	UpdateTeam(ctx context.Context, p *core.Team) error
	SelectAllTeams(ctx context.Context) ([]core.Team, error)
	SelectUserTeams(ctx context.Context, u *core.User) ([]core.Team, error)
	DeleteTeam(ctx context.Context, p *core.Team) error
	DeleteUserTeam(ctx context.Context, u *core.User, p *core.Team) error
	InsertUserTeams(ctx context.Context, userID, TeamID int) error
	SelectPossibleNewUserTeams(ctx context.Context, u *core.User) ([]core.Team, error)
	SelectAllTeamEpics(ctx context.Context, c *core.Team) ([]core.Epic, error)
}
