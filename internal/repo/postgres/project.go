package postgres

import (
	"context"
	"money/internal/core"
	"money/internal/logger"

	"go.uber.org/zap"
)

func (s *postgresStorage) InsertProject(ctx context.Context, p *core.Project) error {
	_, err := s.preparedStatements["insertProject"].ExecContext(ctx, p.Name)
	if err != nil {
		return err
	}
	return nil
}

func (s *postgresStorage) UpdateProject(ctx context.Context, p *core.Project) error {
	//SET Name = $2 WHERE ID = $1
	_, err := s.preparedStatements["updateProject"].ExecContext(ctx, p.ID, p.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s *postgresStorage) SelectProject(ctx context.Context, p *core.Project) error {
	return s.preparedStatements["selectProject"].GetContext(ctx, p, p.ID)
}

func (s *postgresStorage) SelectAllProjects(ctx context.Context) ([]core.Project, error) {
	p := make([]core.Project, 0)
	err := s.preparedStatements["selectAllProjects"].SelectContext(ctx, &p)
	if err != nil {
		logger.Log.Error("selectAllProjects",
			zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (s *postgresStorage) SelectUserProjects(ctx context.Context, u *core.User) ([]core.Project, error) {
	p := make([]core.Project, 0)
	err := s.preparedStatements["selectUserProjects"].SelectContext(ctx, &p, u.ID)
	if err != nil {
		logger.Log.Error("selectAllProjects",
			zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (s *postgresStorage) DeleteProject(ctx context.Context, p *core.Project) error {
	_, err := s.preparedStatements["deleteProject"].ExecContext(ctx, p.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *postgresStorage) DeleteUserProject(ctx context.Context, u *core.User, p *core.Project) error {
	_, err := s.preparedStatements["deleteProject"].ExecContext(ctx, u.ID, p.ID)
	if err != nil {
		logger.Log.Error("DeleteUserProject",
			zap.Error(err))
		return err
	}

	return nil
}

func (s *postgresStorage) InsertUserProject(ctx context.Context, userID, projectID int) error {
	_, err := s.preparedStatements["insertUserProject"].ExecContext(ctx, userID, projectID)
	if err != nil {
		logger.Log.Error("InsertUserProject",
			zap.Error(err))
		return err
	}

	return nil
}

func (s *postgresStorage) SelectPossibleNewUserProjects(ctx context.Context, u *core.User) ([]core.Project, error) {
	var ps []core.Project
	err := s.preparedStatements["selectPossibleUserProjects"].SelectContext(ctx, &ps, u.ID)
	if err != nil {
		logger.Log.Error("selectPossibleUserProjects",
			zap.Error(err))
		return nil, err
	}

	return ps, nil
}
