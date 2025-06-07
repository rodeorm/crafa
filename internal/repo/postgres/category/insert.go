package category

import (
	"context"
	"crafa/internal/core"
)

func (s *Storage) InsertCategory(ctx context.Context, c *core.Category) error {
	return s.stmt["insertCategory"].GetContext(ctx, c, c.Name, c.Level.ID)
}
