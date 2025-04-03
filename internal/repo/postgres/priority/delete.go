package priority

import (
	"context"
	"money/internal/core"
)

func (s *Storage) DeletePriority(ctx context.Context, c *core.Priority) error {
	_, err := s.stmt["deletePriority"].ExecContext(ctx, c.ID)
	return err
}
