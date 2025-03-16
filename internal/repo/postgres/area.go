package postgres

import (
	"context"
	"money/internal/core"

	"github.com/pkg/errors"
)

func (s *postgresStorage) areaPrepareStmts() error {
	insertArea, err := s.DB.Preparex(`INSERT INTO ref.Areas
									(levelid, teamid, name) 
	 								SELECT $1, $2, $3
	 								RETURNING id;`)
	if err != nil {
		return errors.Wrap(err, "insertArea")
	}
	s.preparedStatements["insertArea"] = insertArea
	return nil
}

func (s *postgresStorage) AddArea(ctx context.Context, a *core.Area) error {
	return nil
}
func (s *postgresStorage) EditArea(ctx context.Context, a *core.Area) error {
	return nil
}
func (s *postgresStorage) SelectArea(ctx context.Context, a *core.Area) error {
	return nil
}
func (s *postgresStorage) SelectAllAreas(ctx context.Context) ([]core.Area, error) {
	return nil, nil
}

func (s *postgresStorage) SelectAllTeamLevelAreas(ctx context.Context, t *core.Team, a *core.Area) []core.Area {
	return nil
}

func (s *postgresStorage) DeleteArea(ctx context.Context, a *core.Area) error {
	return nil
}
