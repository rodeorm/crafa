package area

import (
	"context"
	"money/internal/core"

	"github.com/pkg/errors"
)

func (s Storage) SelectAllAreas(ctx context.Context) ([]core.Area, error) {
	a := make([]core.Area, 0)
	err := s.stmt["SelectAllAreas"].SelectContext(ctx, &a)
	if err != nil {
		return nil, errors.Wrap(err, "SelectAllAreas")
	}
	return a, nil
}
