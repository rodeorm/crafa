package priority

import (
	"context"
	"money/internal/core"
)

func (s *Storage) InsertPriority(ctx context.Context, c *core.Priority) error {
	return s.stmt["insertPriority"].GetContext(ctx, c, c.Name, c.Level.ID)
}
