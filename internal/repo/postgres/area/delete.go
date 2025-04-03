package area

import (
	"context"
	"money/internal/core"

	"github.com/pkg/errors"
)

func (s Storage) DeleteArea(ctx context.Context, a *core.Area) error {
	_, err := s.stmt["DeleteArea"].ExecContext(ctx, a.ID)
	if err != nil {
		return errors.Wrap(err, "DeleteArea")
	}
	return nil
}
