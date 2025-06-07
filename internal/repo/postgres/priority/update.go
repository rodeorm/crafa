package priority

import (
	"context"
	"crafa/internal/core"
)

func (s *Storage) UpdatePriority(ctx context.Context, c *core.Priority) error {
	_, err := s.stmt["updatePriority"].ExecContext(ctx, c.ID, c.Name, c.Level.ID)
	return err
}
