package area

import (
	"context"
	"crafa/internal/core"

	"github.com/pkg/errors"
)

func (s Storage) UpdateArea(ctx context.Context, a *core.Area) error {
	_, err := s.stmt["updateArea"].ExecContext(ctx, a.ID, a.Name, a.Level.ID)
	if err != nil {
		return errors.Wrap(err, "SelectAllAreas")
	}
	return nil
}
