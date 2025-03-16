package core

import (
	"context"
)

type Team struct {
	Name string
	ID   int
}

type TeamStorager interface {
	InsertTeam(ctx context.Context, p *Team) error
	SelectTeam(ctx context.Context, p *Team) error
	UpdateTeam(ctx context.Context, p *Team) error
	SelectAllTeams(ctx context.Context) ([]Team, error)
	SelectUserTeams(ctx context.Context, u *User) ([]Team, error)
	DeleteTeam(ctx context.Context, p *Team) error
	DeleteUserTeam(ctx context.Context, u *User, p *Team) error
	InsertUserTeams(ctx context.Context, userID, TeamID int) error
	SelectPossibleNewUserTeams(ctx context.Context, u *User) ([]Team, error)
	SelectAllTeamEpics(ctx context.Context, c *Team) ([]Epic, error)
}
