package postgres

import (
	"context"
	"money/internal/core"
	"money/internal/logger"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (s *postgresStorage) iterationPrepareStmts() error {
	insertIteration, err := s.DB.Preparex(`		INSERT INTO ref.Iterations
												(Name, LevelID, ParentID, Year, Month)
												SELECT $1, $2, $3, $4, $5
												RETURNING id;`)
	if err != nil {
		return errors.Wrap(err, "insertIteration")
	}

	updateIteration, err := s.DB.Preparex(`		UPDATE ref.Iterations
												SET Name = $2, LevelID = $3, ParentID = $4, Year = $5, Month = $6
												WHERE ID = $1;`)
	if err != nil {
		return errors.Wrap(err, "updateIteration")
	}

	selectIteration, err := s.DB.Preparex(`		SELECT i.id, i.name, i.year, i.month,
														l.id AS "level.id", l.name AS "level.name", 
														COALESCE(p.id,0) AS "parent.id", COALESCE(p.name,'-') AS "parent.name"
												FROM ref.Iterations AS i
													LEFT JOIN ref.Iterations AS p ON p.ID = i.ParentID
													INNER JOIN ref.Levels AS l ON l.ID = i.levelID
												WHERE i.ID = $1;`)
	if err != nil {
		return errors.Wrap(err, "selectIteration")
	}

	selectAllIterations, err := s.DB.Preparex(`			SELECT  i.id, i.name, i.year, i.month,
															l.id AS "level.id", l.name AS "level.name", 
															COALESCE(p.id,0) AS "parent.id", COALESCE(p.name,'-') AS "parent.name"
														FROM ref.Iterations AS i
															LEFT JOIN ref.Iterations AS p ON p.ID = i.ParentID
														INNER JOIN ref.Levels AS l ON l.ID = i.levelID;`)
	if err != nil {
		return errors.Wrap(err, "selectAllIterations")
	}

	deleteIteration, err := s.DB.Preparex(`	DELETE FROM ref.Iterations
											WHERE ID = $1;`)
	if err != nil {
		return errors.Wrap(err, "deleteIteration")
	}

	selectPossibleLevelIterations, err := s.DB.Preparex(`			SELECT i.id, i.name, i.year, i.month,
																		l.id AS "level.id", l.name AS "level.name", 
																		COALESCE(p.id,0) AS "parent.id", COALESCE(p.name,'-') AS "parent.name"
																	FROM ref.Iterations AS i
																		LEFT JOIN ref.Iterations AS p ON p.ID = i.ParentID
																		INNER JOIN ref.Levels AS l ON l.ID = i.levelID
																	WHERE i.LevelID = $1;`)
	if err != nil {
		return errors.Wrap(err, "selectPossibleLevelIterations")
	}

	selectPossibleParentIterations, err := s.DB.Preparex(`			SELECT 	p.id AS "id", p.name AS "name",
																			l.id AS "level.id", l.name AS "level.name"
																	FROM ref.Iterations AS i
																		INNER JOIN ref.Iterations AS p ON p.LevelID = i.LevelID - 1
																		INNER JOIN ref.Levels AS l ON l.ID = p.LevelID
																	WHERE i.ID = $1;`)
	if err != nil {
		return errors.Wrap(err, "selectPossibleParentIterations")
	}

	s.preparedStatements["insertIteration"] = insertIteration
	s.preparedStatements["updateIteration"] = updateIteration
	s.preparedStatements["selectIteration"] = selectIteration
	s.preparedStatements["selectAllIterations"] = selectAllIterations
	s.preparedStatements["deleteIteration"] = deleteIteration
	s.preparedStatements["selectPossibleLevelIterations"] = selectPossibleLevelIterations
	s.preparedStatements["selectPossibleParentIterations"] = selectPossibleParentIterations
	return nil
}

func (s *postgresStorage) InsertIteration(ctx context.Context, p *core.Iteration) error {
	//Name, LevelID, ParentID, Year, Month)
	_, err := s.preparedStatements["insertIteration"].ExecContext(ctx, p.Name, p.Level.ID, p.Parent.ID, p.Year, p.Month)
	if err != nil {
		return err
	}
	return nil
}

func (s *postgresStorage) UpdateIteration(ctx context.Context, p *core.Iteration) error {
	_, err := s.preparedStatements["updateIteration"].ExecContext(ctx, p.ID, p.Name, p.Level.ID, p.Parent.ID, p.Year, p.Month)
	if err != nil {
		return err
	}

	return nil
}

func (s *postgresStorage) SelectIteration(ctx context.Context, p *core.Iteration) error {
	return s.preparedStatements["selectIteration"].GetContext(ctx, p, p.ID)
}

func (s *postgresStorage) SelectAllIterations(ctx context.Context) ([]core.Iteration, error) {
	p := make([]core.Iteration, 0)
	err := s.preparedStatements["selectAllIterations"].SelectContext(ctx, &p)
	if err != nil {
		logger.Log.Error("selectAllIterations",
			zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (s *postgresStorage) SelectPossibleLevelIterations(ctx context.Context, l *core.Level) ([]core.Iteration, error) {
	p := make([]core.Iteration, 0)
	err := s.preparedStatements["selectPossibleLevelIterations"].SelectContext(ctx, &p, l.ID)
	if err != nil {
		logger.Log.Error("selectPossibleLevelIterations",
			zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (s *postgresStorage) SelectPossibleParentIterations(ctx context.Context, i *core.Iteration) ([]core.Iteration, error) {
	p := make([]core.Iteration, 0)
	err := s.preparedStatements["selectPossibleLevelIterations"].SelectContext(ctx, &p, i.ID)
	if err != nil {
		logger.Log.Error("selectPossibleLevelIterations",
			zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (s *postgresStorage) DeleteIteration(ctx context.Context, p *core.Iteration) error {
	_, err := s.preparedStatements["deleteIteration"].ExecContext(ctx, p.ID)
	if err != nil {
		return err
	}

	return nil
}
