package iteration

import (
	"context"
	"crafa/internal/core"
	"crafa/internal/logger"

	"go.uber.org/zap"
)

func (s *Storage) SelectPossibleParentIterations(ctx context.Context, i *core.Iteration) ([]core.Iteration, error) {
	p := make([]core.Iteration, 0)
	err := s.stmt["selectPossibleLevelIterations"].SelectContext(ctx, &p, i.ID)
	if err != nil {
		logger.Log.Error("selectPossibleLevelIterations",
			zap.Error(err))
		return nil, err
	}
	return p, nil
}
