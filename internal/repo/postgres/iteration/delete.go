package iteration

import (
	"context"
	"crafa/internal/core"
)

func (s *Storage) DeleteIteration(ctx context.Context, p *core.Iteration) error {
	_, err := s.stmt["deleteIteration"].ExecContext(ctx, p.ID)
	if err != nil {
		return err
	}

	return nil
}
