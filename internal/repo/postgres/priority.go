package postgres

import (
	"context"
	"money/internal/core"
	"money/internal/logger"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (s *postgresStorage) InsertPriority(ctx context.Context, c *core.Priority) error {
	return s.preparedStatements["insertPriority"].GetContext(ctx, c, c.Name, c.Level.ID)
}
func (s *postgresStorage) UpdatePriority(ctx context.Context, c *core.Priority) error {
	_, err := s.preparedStatements["updatePriority"].ExecContext(ctx, c.ID, c.Name, c.Level.ID)
	return err
}
func (s *postgresStorage) SelectPriority(ctx context.Context, c *core.Priority) error {
	return s.preparedStatements["selectPriority"].GetContext(ctx, c, c.ID)
}

func (s *postgresStorage) SelectAllPriorities(ctx context.Context) ([]core.Priority, error) {
	c := make([]core.Priority, 0)
	err := s.preparedStatements["selectAllPriorities"].SelectContext(ctx, &c)
	if err != nil {
		logger.Log.Error("SelectAllPriorities",
			zap.Error(err))
		return nil, err
	}
	return c, nil
}

func (s *postgresStorage) SelectAllLevelPriorities(ctx context.Context, l *core.Level) error {
	var c []core.Priority
	err := s.preparedStatements["selectLevelPriorities"].SelectContext(ctx, c, l.ID)
	if err != nil {
		return err
	}
	l.PossiblePriorities = c
	return nil
}

func (s *postgresStorage) DeletePriority(ctx context.Context, c *core.Priority) error {
	_, err := s.preparedStatements["deletePriority"].ExecContext(ctx, c.ID)
	return err
}

func (s *postgresStorage) priorityPrepareStmts() error {
	insertPriority, err := s.DB.Preparex(`		INSERT INTO ref.Priorities
												(Name, LevelID)
												SELECT $1, $2
												RETURNING id;`)
	if err != nil {
		return errors.Wrap(err, "insertPriority")
	}

	updatePriority, err := s.DB.Preparex(`		UPDATE ref.Priorities
												SET Name = $2, LevelID = $3
												WHERE ID = $1;`)
	if err != nil {
		return errors.Wrap(err, "updatePriority")
	}

	selectPriority, err := s.DB.Preparex(`		SELECT c.id, c.name, l.ID AS "level.id", l.name AS "level.name"
												FROM ref.Priorities AS c
													INNER JOIN ref.Levels AS l ON l.ID = c.LevelID
												WHERE c.ID = $1;`)
	if err != nil {
		return errors.Wrap(err, "selectPriority")
	}

	selectAllPriorities, err := s.DB.Preparex(`		SELECT c.id as "id", c.name as "name", l.id AS "level.id", l.name AS "level.name"
													FROM ref.Priorities AS c 
														INNER JOIN ref.Levels AS l ON l.ID = c.LevelID ;`)
	if err != nil {
		return errors.Wrap(err, "selectAllPriorities")
	}

	selectLevelPriorities, err := s.DB.Preparex(`	SELECT c.id as "id", c.name as "name", l.id AS "level.id", l.name AS "level.name"
													FROM ref.Priorities AS c 
														INNER JOIN ref.Levels AS l ON l.ID = c.LevelID 
													WHERE levelid = $1;`)
	if err != nil {
		return errors.Wrap(err, "selectLevelPriorities")
	}

	deletePriority, err := s.DB.Preparex(`	DELETE FROM ref.Priorities
											WHERE ID = $1;`)
	if err != nil {
		return errors.Wrap(err, "deletePriority")
	}

	s.preparedStatements["insertPriority"] = insertPriority
	s.preparedStatements["updatePriority"] = updatePriority
	s.preparedStatements["selectPriority"] = selectPriority
	s.preparedStatements["selectAllPriorities"] = selectAllPriorities
	s.preparedStatements["deletePriority"] = deletePriority
	s.preparedStatements["selectLevelPriorities"] = selectLevelPriorities

	return nil
}
