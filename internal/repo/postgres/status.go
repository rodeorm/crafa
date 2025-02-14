package postgres

import (
	"context"
	"money/internal/core"
)

func (s *postgresStorage) AddStatus(ctx context.Context, a *core.Status) error {
	return nil
}
func (s *postgresStorage) EditStatus(ctx context.Context, a *core.Status) error {
	return nil
}
func (s *postgresStorage) SelectStatus(ctx context.Context, a *core.Status) error {
	return nil
}
func (s *postgresStorage) SelectAllStatuses(ctx context.Context) ([]core.Status, error) {
	return nil, nil
}

func (s *postgresStorage) SelectAllLevelStatuses(ctx context.Context, l *core.Level) []core.Status {
	return nil
}

func (s *postgresStorage) DeleteStatus(ctx context.Context, a *core.Status) error {
	return nil
}
