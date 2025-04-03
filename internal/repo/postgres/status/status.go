package status

import (
	"context"
	"money/internal/core"

	"github.com/pkg/errors"
)

func (s *Storage) InsertStatus(ctx context.Context, a *core.Status) error {
	return s.stmt["insertStatus"].GetContext(ctx, a, a.Name, a.Level.ID)
}

func (s *Storage) InsertStatusHierarchy(ctx context.Context, parentID, childID int) error {
	_, err := s.stmt["insertStatusHierarchy"].ExecContext(ctx, parentID, childID)
	return err
}

func (s *Storage) UpdateStatus(ctx context.Context, a *core.Status) error {
	_, err := s.stmt["updateStatus"].ExecContext(ctx, a.ID, a.Name, a.Level.ID)
	return err
}
func (s *Storage) SelectStatus(ctx context.Context, st *core.Status) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "SelectStatus 1")
	}

	st.Children = make([]core.Status, 0)
	st.Parents = make([]core.Status, 0)

	// Получаем основные данные статуса
	err = tx.Stmtx(s.stmt["selectStatus"]).GetContext(ctx, st, st.ID)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "SelectStatus 2")
	}
	// Получаем родительские статусы
	err = tx.Stmtx(s.stmt["selectParents"]).GetContext(ctx, &st.Parents, st.ID)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "SelectStatus 3")
	}
	// Получаем дочерние статусы
	err = tx.Stmtx(s.stmt["selectChildren"]).GetContext(ctx, &st.Children, st.ID)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "SelectStatus 4")
	}
	return tx.Commit()
}

func (s *Storage) SelectAllStatuses(ctx context.Context) ([]core.Status, error) {
	a := make([]core.Status, 0)
	err := s.stmt["selectAllStatuses"].SelectContext(ctx, &a)
	if err != nil {

		return nil, err
	}
	return a, nil
}

func (s *Storage) SelectAllLevelStatuses(ctx context.Context, l *core.Level) error {
	var c []core.Status
	err := s.stmt["selectLevelStatuses"].SelectContext(ctx, c, l.ID)
	if err != nil {
		return err
	}
	l.PossibleStatuses = c
	return nil
}

func (s *Storage) SelectFirstLevelStatuses(ctx context.Context, l *core.Level) error {
	var c []core.Status
	err := s.stmt["selectFirstLevelStatuses"].SelectContext(ctx, c, l.ID)
	if err != nil {
		return err
	}
	l.PossibleStatuses = c
	return nil
}

func (s *Storage) SelectPossibleParents(ctx context.Context, st *core.Status) ([]core.Status, error) {
	var parents []core.Status
	err := s.stmt["selectStatusPossibleParents"].SelectContext(ctx, parents, st.ID)
	if err != nil {
		return nil, err
	}
	return parents, nil
}

func (s *Storage) SelectPossibleChildren(ctx context.Context, st *core.Status) ([]core.Status, error) {
	var parents []core.Status
	err := s.stmt["selectStatusPossibleChildren"].SelectContext(ctx, parents, st.ID)
	if err != nil {
		return nil, err
	}
	return parents, nil
}

func (s *Storage) DeleteStatus(ctx context.Context, a *core.Status) error {
	_, err := s.stmt["deleteStatus"].ExecContext(ctx, a.ID)
	return err
}
