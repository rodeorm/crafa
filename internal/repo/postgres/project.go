package postgres

import (
	"context"
	"log"
	"money/internal/core"
	"money/internal/logger"

	"go.uber.org/zap"
)

func (s *PostgresStorage) InsertProject(ctx context.Context, p *core.Project) error {
	_, err := s.preparedStatements["insertProject"].ExecContext(ctx, p.Name)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStorage) UpdateProject(ctx context.Context, p *core.Project) error {
	//SET Name = $2 WHERE ID = $1
	_, err := s.preparedStatements["updateProject"].ExecContext(ctx, p.ID, p.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStorage) SelectProject(ctx context.Context, p *core.Project) error {

	tx, err := s.DB.BeginTxx(ctx, nil)
	if err != nil {
		log.Println("selectProject 1", err)
		return err
	}
	err = tx.Stmtx(s.preparedStatements["selectProject"]).GetContext(ctx, p, p.ID)
	if err != nil {
		log.Println("selectProject 2", err)
		tx.Rollback()
		return err
	}

	p.Epics = make([]core.Epic, 0)
	err = tx.Stmtx(s.preparedStatements["selectProjectEpics"]).SelectContext(ctx, &p.Epics, p.ID)
	if err != nil {
		log.Println("selectProject 2", err)
		tx.Rollback()
		return err
	}

	return tx.Commit()

}

func (s *PostgresStorage) SelectAllProjects(ctx context.Context) ([]core.Project, error) {
	p := make([]core.Project, 0)
	err := s.preparedStatements["selectAllProjects"].SelectContext(ctx, &p)
	if err != nil {
		logger.Log.Error("selectAllProjects",
			zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (s *PostgresStorage) SelectUserProjects(ctx context.Context, u *core.User) ([]core.Project, error) {
	p := make([]core.Project, 0)
	err := s.preparedStatements["selectUserProjects"].SelectContext(ctx, &p, u.ID)
	if err != nil {
		logger.Log.Error("selectAllProjects",
			zap.Error(err))
		return nil, err

	}
	return p, nil
}

func (s *PostgresStorage) DeleteProject(ctx context.Context, p *core.Project) error {
	_, err := s.preparedStatements["deleteProject"].ExecContext(ctx, p.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStorage) DeleteUserProject(ctx context.Context, u *core.User, p *core.Project) error {
	_, err := s.preparedStatements["deleteUserProject"].ExecContext(ctx, u.ID, p.ID)
	if err != nil {
		logger.Log.Error("DeleteUserProject",
			zap.Error(err))
		return err
	}

	return nil
}

func (s *PostgresStorage) InsertUserProject(ctx context.Context, userID, projectID int) error {
	_, err := s.preparedStatements["insertUserProject"].ExecContext(ctx, userID, projectID)
	if err != nil {
		logger.Log.Error("InsertUserProject",
			zap.Error(err))
		return err
	}

	return nil
}

func (s *PostgresStorage) SelectPossibleNewUserProjects(ctx context.Context, u *core.User) ([]core.Project, error) {
	var ps []core.Project
	err := s.preparedStatements["selectPossibleUserProjects"].SelectContext(ctx, &ps, u.ID)
	if err != nil {
		logger.Log.Error("selectPossibleUserProjects",
			zap.Error(err))
		return nil, err
	}

	return ps, nil
}

func (s *PostgresStorage) SelectAllProjectEpics(ctx context.Context, c *core.Project) ([]core.Epic, error) {
	return nil, nil
}
