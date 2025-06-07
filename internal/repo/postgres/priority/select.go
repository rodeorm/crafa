package priority

import (
	"context"
	"crafa/internal/core"
)

func (s *Storage) SelectPriority(ctx context.Context, c *core.Priority) error {
	return s.stmt["selectPriority"].GetContext(ctx, c, c.ID)
}
