package area

import (
	"context"
	"crafa/internal/core"

	"github.com/pkg/errors"
)

func (s Storage) InsertArea(ctx context.Context, a *core.Area) error {
	err := s.stmt["InsertArea"].GetContext(ctx, a, a.Name, a.Level.ID)
	if err != nil {
		return errors.Wrap(err, "InsertArea")
	}
	return nil
}
