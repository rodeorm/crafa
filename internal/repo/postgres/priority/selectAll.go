package priority

import (
	"context"
	"money/internal/core"
	"money/internal/logger"

	"go.uber.org/zap"
)

func (s *Storage) SelectAllPriorities(ctx context.Context) ([]core.Priority, error) {
	c := make([]core.Priority, 0)
	err := s.stmt["selectAllPriorities"].SelectContext(ctx, &c)
	if err != nil {
		logger.Log.Error("SelectAllPriorities",
			zap.Error(err))
		return nil, err
	}
	return c, nil
}
