package postgres

import (
	"context"
	"money/internal/core"

	"github.com/pkg/errors"
)

func (s *postgresStorage) InsertCategory(ctx context.Context, c *core.Category) error {
	return s.preparedStatements["insertCategory"].GetContext(ctx, c, c.Name, c.Level.ID)
}
func (s *postgresStorage) UpdateCategory(ctx context.Context, c *core.Category) error {
	_, err := s.preparedStatements["updateCategory"].ExecContext(ctx, c.ID, c.Name, c.Level.ID)
	return err
}
func (s *postgresStorage) SelectCategory(ctx context.Context, c *core.Category) error {
	return s.preparedStatements["selectCategory"].GetContext(ctx, c, c.ID)
}

func (s *postgresStorage) SelectAllCategories(ctx context.Context) ([]core.Category, error) {
	c := make([]core.Category, 0)
	err := s.preparedStatements["selectAllCategories"].SelectContext(ctx, &c)
	if err != nil {

		return nil, err
	}
	return c, nil
}

func (s *postgresStorage) SelectAllLevelCategories(ctx context.Context, l *core.Level) error {
	var c []core.Category
	err := s.preparedStatements["selectLevelCategories"].SelectContext(ctx, c, l.ID)
	if err != nil {
		return err
	}
	l.PossibleCategories = c
	return nil
}

func (s *postgresStorage) DeleteCategory(ctx context.Context, c *core.Category) error {
	_, err := s.preparedStatements["deleteCategory"].ExecContext(ctx, c.ID)
	return err
}

func (s *postgresStorage) categoryPrepareStmts() error {
	insertCategory, err := s.DB.Preparex(`		INSERT INTO ref.Categories
												(Name, LevelID)
												SELECT $1, $2
												RETURNING id;`)
	if err != nil {
		return errors.Wrap(err, "insertCategory")
	}

	updateCategory, err := s.DB.Preparex(`		UPDATE ref.Categories
												SET Name = $2, LevelID = $3
												WHERE ID = $1;`)
	if err != nil {
		return errors.Wrap(err, "updateCategory")
	}

	selectCategory, err := s.DB.Preparex(`		SELECT id, name, levelid
												FROM ref.Categories
												WHERE ID = $1;`)
	if err != nil {
		return errors.Wrap(err, "selectCategory")
	}

	selectAllCategories, err := s.DB.Preparex(`		SELECT c.id as "id", c.name as "name", l.id AS "level.id", l.name AS "level.name"
													FROM ref.Categories AS c 
														INNER JOIN ref.Levels AS l ON l.ID = c.LevelID ;`)
	if err != nil {
		return errors.Wrap(err, "selectAllCategories")
	}

	selectLevelCategories, err := s.DB.Preparex(`	SELECT c.id as "id", c.name as "name", l.id AS "level.id", l.name AS "level.name"
													FROM ref.Categories AS c 
														INNER JOIN ref.Levels AS l ON l.ID = c.LevelID 
													WHERE levelid = $1;`)
	if err != nil {
		return errors.Wrap(err, "selectLevelCategories")
	}

	deleteCategory, err := s.DB.Preparex(`	DELETE FROM ref.Categories
											WHERE ID = $1;`)
	if err != nil {
		return errors.Wrap(err, "deleteCategory")
	}

	s.preparedStatements["insertCategory"] = insertCategory
	s.preparedStatements["updateCategory"] = updateCategory
	s.preparedStatements["selectCategory"] = selectCategory
	s.preparedStatements["selectAllCategories"] = selectAllCategories
	s.preparedStatements["deleteCategory"] = deleteCategory
	s.preparedStatements["selectLevelCategories"] = selectLevelCategories

	return nil
}
