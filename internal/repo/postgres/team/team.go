package team

import (
	"context"
	"log"
	"money/internal/core"
	"money/internal/logger"

	"go.uber.org/zap"
)

func (s *Storage) InsertTeam(ctx context.Context, p *core.Team) error {
	_, err := s.stmt["insertTeam"].ExecContext(ctx, p.Name)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) UpdateTeam(ctx context.Context, p *core.Team) error {
	//SET Name = $2 WHERE ID = $1
	_, err := s.stmt["updateTeam"].ExecContext(ctx, p.ID, p.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) SelectTeam(ctx context.Context, t *core.Team) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		log.Println("SelectTeam 1", err)
		return err
	}
	err = tx.Stmtx(s.stmt["selectTeam"]).GetContext(ctx, t, t.ID)
	if err != nil {
		log.Println("SelectTeam 2", err)
		tx.Rollback()
		return err
	}

	t.Users = make([]core.User, 0)
	err = tx.Stmtx(s.stmt["selectTeamUsers"]).SelectContext(ctx, &t.Users, t.ID)
	if err != nil {
		log.Println("SelectTeam 3", err)
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (s *Storage) SelectAllTeams(ctx context.Context) ([]core.Team, error) {
	p := make([]core.Team, 0)
	err := s.stmt["selectAllTeams"].SelectContext(ctx, &p)
	if err != nil {
		logger.Log.Error("selectAllTeams",
			zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (s *Storage) SelectUserTeams(ctx context.Context, u *core.User) ([]core.Team, error) {
	p := make([]core.Team, 0)
	err := s.stmt["selectUserTeams"].SelectContext(ctx, &p, u.ID)
	if err != nil {
		logger.Log.Error("selectAllTeams",
			zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (s *Storage) DeleteTeam(ctx context.Context, p *core.Team) error {
	_, err := s.stmt["deleteTeam"].ExecContext(ctx, p.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) DeleteUserTeam(ctx context.Context, u *core.User, p *core.Team) error {
	_, err := s.stmt["deleteUserTeam"].ExecContext(ctx, u.ID, p.ID)
	if err != nil {
		logger.Log.Error("DeleteUserTeams",
			zap.Error(err))
		return err
	}

	return nil
}

func (s *Storage) InsertUserTeams(ctx context.Context, userID, TeamID int) error {
	_, err := s.stmt["insertUserTeams"].ExecContext(ctx, userID, TeamID)
	if err != nil {
		logger.Log.Error("InsertUserTeams",
			zap.Error(err))
		return err
	}

	return nil
}

func (s *Storage) SelectPossibleNewUserTeams(ctx context.Context, u *core.User) ([]core.Team, error) {
	var ps []core.Team
	err := s.stmt["selectPossibleUserTeams"].SelectContext(ctx, &ps, u.ID)
	if err != nil {
		logger.Log.Error("selectPossibleUserTeams",
			zap.Error(err))
		return nil, err
	}

	return ps, nil
}

func (s *Storage) SelectAllTeamEpics(ctx context.Context, c *core.Team) ([]core.Epic, error) {
	return nil, nil
}
