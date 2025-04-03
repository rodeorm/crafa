package postgres

import (
	"context"
	"money/internal/core"
)

func (s *PostgresStorage) InsertEpic(ctx context.Context, e *core.Epic) error {
	return nil
}
func (s *PostgresStorage) SelectEpic(ctx context.Context, e *core.Epic) error {
	return nil
}
func (s *PostgresStorage) UpdateEpic(ctx context.Context, e *core.Epic) error {
	return nil
}
func (s *PostgresStorage) DeleteEpic(ctx context.Context, e *core.Epic) error {
	return nil
}

func (s *PostgresStorage) SelectAllEpicIssues(ctx context.Context, e *core.Epic) ([]core.Issue, error) {
	return nil, nil
}
func (s *PostgresStorage) SelectAllEpicFeautures(ctx context.Context, e *core.Epic) ([]core.Feature, error) {
	return nil, nil
}
