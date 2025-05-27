package iteration

import (
	"context"
	"money/internal/core"
	"money/internal/logger"

	"go.uber.org/zap"
)

func (s *Storage) SelectAllIterations(ctx context.Context) ([]core.Iteration, error) {
	p := make([]core.Iteration, 0)
	err := s.stmt["selectAllIterations"].SelectContext(ctx, &p)
	if err != nil {
		logger.Log.Error("selectAllIterations",
			zap.Error(err))
		return nil, err
	}
	return p, nil
}
