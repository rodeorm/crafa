package category

import (
	"context"
	"money/internal/core"
)

func (s Storage) SelectAllLevelCategories(ctx context.Context, l *core.Level) error {
	var c []core.Category
	err := s.stmt["selectLevelCategories"].SelectContext(ctx, c, l.ID)
	if err != nil {
		return err
	}
	l.PossibleCategories = c
	return nil
}
