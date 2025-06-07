package area

import (
	"context"
	"crafa/internal/core"

	"github.com/pkg/errors"
)

func (s Storage) SelectAllLevelAreas(ctx context.Context, l *core.Level) error {
	var a []core.Area
	err := s.stmt["SelectAllLevelAreas"].SelectContext(ctx, a, l.ID)
	if err != nil {
		return errors.Wrap(err, "SelectAllLevelAreas")
	}
	l.PossibleAreas = a
	return nil
}
