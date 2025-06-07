package iteration

import (
	"context"
	"crafa/internal/core"
)

func (s *Storage) UpdateIteration(ctx context.Context, p *core.Iteration) error {
	_, err := s.stmt["updateIteration"].ExecContext(ctx, p.ID, p.Name, p.Level.ID, p.Parent.ID, p.Year, p.Month)
	if err != nil {
		return err
	}

	return nil
}
