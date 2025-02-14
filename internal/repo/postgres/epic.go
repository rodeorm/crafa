package postgres

import (
	"context"
	"money/internal/core"
)

func (s *postgresStorage) InsertEpic(ctx context.Context, e *core.Epic) error {
	return nil
}
func (s *postgresStorage) SelectEpic(ctx context.Context, e *core.Epic) error {
	return nil
}
func (s *postgresStorage) UpdateEpic(ctx context.Context, e *core.Epic) error {
	return nil
}
func (s *postgresStorage) DeleteEpic(ctx context.Context, e *core.Epic) error {
	return nil
}

func (s *postgresStorage) SelectAllEpicIssues(ctx context.Context, e *core.Epic) ([]core.Issue, error) {
	return nil, nil
}
func (s *postgresStorage) SelectAllEpicFeautures(ctx context.Context, e *core.Epic) ([]core.Feature, error) {
	return nil, nil
}
