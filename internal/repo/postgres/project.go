package postgres

import (
	"context"
	"money/internal/core"
)

func (s *postgresStorage) AddProject(ctx context.Context, p *core.Project, u *core.User) error {
	return nil
}

func (s *postgresStorage) EditProject(ctx context.Context, p *core.Project, u *core.User) error {
	return nil
}

func (s *postgresStorage) SelectProject(ctx context.Context, p *core.Project, u *core.User) error {
	return nil
}

func (s *postgresStorage) SelectAllProjects(ctx context.Context, u *core.User) ([]core.Project, error) {
	return nil, nil
}

func (s *postgresStorage) DeleteProject(ctx context.Context, p *core.Project, u *core.User) error {
	return nil
}
