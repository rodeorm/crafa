package category

import (
	"context"
	"crafa/internal/core"
)

func (s Storage) SelectCategory(ctx context.Context, c *core.Category) error {
	return s.stmt["selectCategory"].GetContext(ctx, c, c.ID)
}
