package project

import (
	"context"
	"database/sql"
	"log"
	"crafa/internal/core"
	"crafa/internal/logger"

	"go.uber.org/zap"
)

func (s *Storage) InsertProject(ctx context.Context, p *core.Project) error {
	_, err := s.stmt["insertProject"].ExecContext(ctx, p.Name)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) UpdateProject(ctx context.Context, p *core.Project) error {
	//SET Name = $2 WHERE ID = $1
	_, err := s.stmt["updateProject"].ExecContext(ctx, p.ID, p.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) SelectProject(ctx context.Context, p *core.Project) error {

	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		log.Println("selectProject 1", err)
		return err
	}
	err = tx.Stmtx(s.stmt["selectProject"]).GetContext(ctx, p, p.ID)
	if err != nil {
		log.Println("selectProject 2", err)
		tx.Rollback()
		return err
	}

	p.Epics = make([]core.Epic, 0)
	err = tx.Stmtx(s.stmt["selectProjectEpics"]).SelectContext(ctx, &p.Epics, p.ID)
	if err != nil && err != sql.ErrNoRows {
		log.Println("selectProject 2", err)
		tx.Rollback()
		return err
	}

	return tx.Commit()

}

func (s *Storage) SelectAllProjects(ctx context.Context) ([]core.Project, error) {
	p := make([]core.Project, 0)
	err := s.stmt["selectAllProjects"].SelectContext(ctx, &p)
	if err != nil {
		logger.Log.Error("selectAllProjects",
			zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (s *Storage) SelectUserProjects(ctx context.Context, u *core.User) ([]core.Project, error) {
	p := make([]core.Project, 0)
	err := s.stmt["selectUserProjects"].SelectContext(ctx, &p, u.ID)
	if err != nil {
		logger.Log.Error("selectAllProjects",
			zap.Error(err))
		return nil, err

	}
	return p, nil
}

func (s *Storage) DeleteProject(ctx context.Context, p *core.Project) error {
	_, err := s.stmt["deleteProject"].ExecContext(ctx, p.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) DeleteUserProject(ctx context.Context, u *core.User, p *core.Project) error {
	_, err := s.stmt["deleteUserProject"].ExecContext(ctx, u.ID, p.ID)
	if err != nil {
		logger.Log.Error("DeleteUserProject",
			zap.Error(err))
		return err
	}

	return nil
}

func (s *Storage) InsertUserProject(ctx context.Context, userID, projectID int) error {
	_, err := s.stmt["insertUserProject"].ExecContext(ctx, userID, projectID)
	if err != nil {
		logger.Log.Error("InsertUserProject",
			zap.Error(err))
		return err
	}

	return nil
}

func (s *Storage) SelectPossibleNewUserProjects(ctx context.Context, u *core.User) ([]core.Project, error) {
	var ps []core.Project
	err := s.stmt["selectPossibleUserProjects"].SelectContext(ctx, &ps, u.ID)
	if err != nil {
		logger.Log.Error("selectPossibleUserProjects",
			zap.Error(err))
		return nil, err
	}

	return ps, nil
}

func (s *Storage) SelectAllProjectEpics(ctx context.Context, c *core.Project) ([]core.Epic, error) {
	return nil, nil
}
