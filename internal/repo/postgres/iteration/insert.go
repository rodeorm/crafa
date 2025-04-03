package iteration

import (
	"context"
	"money/internal/core"
)

func (s *Storage) InsertIteration(ctx context.Context, i *core.Iteration) error {
	err := s.stmt["insertIteration"].GetContext(ctx, &i.ID, i.Name, i.Level.ID, 0, i.Year, i.Month)
	if err != nil {
		return err
	}
	return nil
}
