package priority

import (
	"context"
	"money/internal/core"
)

func (s *Storage) SelectAllLevelPriorities(ctx context.Context, l *core.Level) error {
	var c []core.Priority
	err := s.stmt["selectLevelPriorities"].SelectContext(ctx, c, l.ID)
	if err != nil {
		return err
	}
	l.PossiblePriorities = c
	return nil
}
