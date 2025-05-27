package area

import (
	"context"
	"money/internal/core"

	"github.com/pkg/errors"
)

func (s Storage) SelectArea(ctx context.Context, a *core.Area) error {
	err := s.stmt["SelectArea"].GetContext(ctx, a, a.ID)
	if err != nil {
		return errors.Wrap(err, "InsertArea")
	}
	return nil
}
