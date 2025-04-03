package category

import (
	"context"
	"money/internal/core"
)

func (s Storage) SelectAllCategories(ctx context.Context) ([]core.Category, error) {
	c := make([]core.Category, 0)
	err := s.stmt["selectAllCategories"].SelectContext(ctx, &c)
	if err != nil {

		return nil, err
	}
	return c, nil
}
