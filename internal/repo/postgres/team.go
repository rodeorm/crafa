package postgres

import (
	"context"
	"log"
	"money/internal/core"
	"money/internal/logger"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (s *postgresStorage) teamPrepareStmts() error {
	insertTeam, err := s.DB.Preparex(`		INSERT INTO cmn.Teams
												(Name)
												SELECT $1
												RETURNING id;`)
	if err != nil {
		return errors.Wrap(err, "insertTeam")
	}

	insertUserTeams, err := s.DB.Preparex(`	INSERT INTO cmn.UserTeams
	(UserID, TeamID)
	SELECT $1, $2;`)
	if err != nil {
		return errors.Wrap(err, "insertUserTeams")
	}

	updateTeam, err := s.DB.Preparex(`			UPDATE cmn.Teams
												SET Name = $2
												WHERE ID = $1;`)
	if err != nil {
		return errors.Wrap(err, "updateTeam")
	}

	selectTeam, err := s.DB.Preparex(`			SELECT id, name
												FROM cmn.Teams
												WHERE ID = $1;`)
	if err != nil {
		return errors.Wrap(err, "selectTeam")
	}

	selectAllTeams, err := s.DB.Preparex(`			SELECT id, name
													FROM cmn.Teams;`)
	if err != nil {
		return errors.Wrap(err, "selectAllTeams")
	}

	selectUserTeams, err := s.DB.Preparex(`		SELECT p.id, p.name
													FROM cmn.Teams AS p
														INNER JOIN cmn.UserTeams AS up
															ON p.ID = up.TeamID
													WHERE up.UserID = $1
													;`)
	if err != nil {
		return errors.Wrap(err, "selectUserTeams")
	}

	deleteTeam, err := s.DB.Preparex(`	DELETE FROM cmn.Teams
											WHERE ID = $1;`)
	if err != nil {
		return errors.Wrap(err, "deleteTeam")
	}

	deleteUserTeam, err := s.DB.Preparex(`	DELETE FROM cmn.UserTeams
												WHERE UserID = $1 AND TeamID = $2;`)
	if err != nil {
		return errors.Wrap(err, "deleteUserTeams")
	}

	selectPossibleUserTeams, err := s.DB.Preparex(`			SELECT p.id, p.name
															FROM cmn.Teams AS p
															LEFT JOIN cmn.UserTeams AS up
															ON p.ID = up.TeamID AND up.UserID = $1
															WHERE up.ID IS NULL
															;`)
	if err != nil {
		return errors.Wrap(err, "selectPossibleUserTeams")
	}

	selectTeamUsers, err := s.DB.Preparex(`			SELECT u.ID AS "user.id", u.Login, u.Name, u.FamilyName, u.PatronName, u.Email
													FROM cmn.UserTeams AS ut
															INNER JOIN cmn.Users AS u 
																ON u.ID  = ut.UserID
													WHERE ut.TeamID = $1;`)
	if err != nil {
		return errors.Wrap(err, "selectTeamUsers")
	}

	s.preparedStatements["insertTeam"] = insertTeam
	s.preparedStatements["updateTeam"] = updateTeam
	s.preparedStatements["selectTeam"] = selectTeam
	s.preparedStatements["selectAllTeams"] = selectAllTeams
	s.preparedStatements["deleteTeam"] = deleteTeam
	s.preparedStatements["selectUserTeams"] = selectUserTeams
	s.preparedStatements["deleteUserTeam"] = deleteUserTeam
	s.preparedStatements["selectPossibleUserTeams"] = selectPossibleUserTeams
	s.preparedStatements["insertUserTeams"] = insertUserTeams
	s.preparedStatements["selectTeamUsers"] = selectTeamUsers

	return nil
}

func (s *postgresStorage) InsertTeam(ctx context.Context, p *core.Team) error {
	_, err := s.preparedStatements["insertTeam"].ExecContext(ctx, p.Name)
	if err != nil {
		return err
	}
	return nil
}

func (s *postgresStorage) UpdateTeam(ctx context.Context, p *core.Team) error {
	//SET Name = $2 WHERE ID = $1
	_, err := s.preparedStatements["updateTeam"].ExecContext(ctx, p.ID, p.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s *postgresStorage) SelectTeam(ctx context.Context, t *core.Team) error {
	tx, err := s.DB.BeginTxx(ctx, nil)
	if err != nil {
		log.Println("SelectTeam 1", err)
		return err
	}
	err = tx.Stmtx(s.preparedStatements["selectTeam"]).GetContext(ctx, t, t.ID)
	if err != nil {
		log.Println("SelectTeam 2", err)
		tx.Rollback()
		return err
	}

	teamUsers := make([]core.User, 0)
	err = tx.Stmtx(s.preparedStatements["selectTeamUsers"]).SelectContext(ctx, &teamUsers, t.ID)
	if err != nil {
		log.Println("SelectTeam 3", err)
		tx.Rollback()
		return err
	}
	t.Users = teamUsers
	tx.Commit()

	return nil
}

func (s *postgresStorage) SelectAllTeams(ctx context.Context) ([]core.Team, error) {
	p := make([]core.Team, 0)
	err := s.preparedStatements["selectAllTeams"].SelectContext(ctx, &p)
	if err != nil {
		logger.Log.Error("selectAllTeams",
			zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (s *postgresStorage) SelectUserTeams(ctx context.Context, u *core.User) ([]core.Team, error) {
	p := make([]core.Team, 0)
	err := s.preparedStatements["selectUserTeams"].SelectContext(ctx, &p, u.ID)
	if err != nil {
		logger.Log.Error("selectAllTeams",
			zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (s *postgresStorage) DeleteTeam(ctx context.Context, p *core.Team) error {
	_, err := s.preparedStatements["deleteTeam"].ExecContext(ctx, p.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *postgresStorage) DeleteUserTeam(ctx context.Context, u *core.User, p *core.Team) error {
	_, err := s.preparedStatements["deleteUserTeam"].ExecContext(ctx, u.ID, p.ID)
	if err != nil {
		logger.Log.Error("DeleteUserTeams",
			zap.Error(err))
		return err
	}

	return nil
}

func (s *postgresStorage) InsertUserTeams(ctx context.Context, userID, TeamID int) error {
	_, err := s.preparedStatements["insertUserTeams"].ExecContext(ctx, userID, TeamID)
	if err != nil {
		logger.Log.Error("InsertUserTeams",
			zap.Error(err))
		return err
	}

	return nil
}

func (s *postgresStorage) SelectPossibleNewUserTeams(ctx context.Context, u *core.User) ([]core.Team, error) {
	var ps []core.Team
	err := s.preparedStatements["selectPossibleUserTeams"].SelectContext(ctx, &ps, u.ID)
	if err != nil {
		logger.Log.Error("selectPossibleUserTeams",
			zap.Error(err))
		return nil, err
	}

	return ps, nil
}

func (s *postgresStorage) SelectAllTeamEpics(ctx context.Context, c *core.Team) ([]core.Epic, error) {
	return nil, nil
}
