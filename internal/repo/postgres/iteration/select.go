package iteration

import (
	"context"
	"money/internal/core"
)

func (s *Storage) SelectIteration(ctx context.Context, p *core.Iteration) error {
	return s.stmt["selectIteration"].GetContext(ctx, p, p.ID)
}
