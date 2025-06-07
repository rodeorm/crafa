package category

import (
	"context"
	"crafa/internal/core"
)

func (s Storage) UpdateCategory(ctx context.Context, c *core.Category) error {
	_, err := s.stmt["updateCategory"].ExecContext(ctx, c.ID, c.Name, c.Level.ID)
	return err
}

func (s Storage) DeleteCategory(ctx context.Context, c *core.Category) error {
	_, err := s.stmt["updateCategory"].ExecContext(ctx, c.ID, c.Name, c.Level.ID)
	return err
}
