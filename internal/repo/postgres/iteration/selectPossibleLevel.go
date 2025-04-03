package iteration

import (
	"context"
	"money/internal/core"
	"money/internal/logger"

	"go.uber.org/zap"
)

func (s *Storage) SelectPossibleLevelIterations(ctx context.Context, l *core.Level) ([]core.Iteration, error) {
	p := make([]core.Iteration, 0)
	err := s.stmt["selectPossibleLevelIterations"].SelectContext(ctx, &p, l.ID)
	if err != nil {
		logger.Log.Error("selectPossibleLevelIterations",
			zap.Error(err))
		return nil, err
	}
	return p, nil
}
