package project

import (
	"context"
	"database/sql"
	"log"
	"money/internal/core"
)

func (s *Storage) SelectUserProject(ctx context.Context, p *core.Project, u *core.User) error {

	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		log.Println("selectProject 1", err)
		return err
	}
	err = tx.Stmtx(s.stmt["selectUserProject"]).GetContext(ctx, p, p.ID, u.ID)
	if err != nil {
		log.Println("selectProject 2", err)
		tx.Rollback()
		return err
	}

	p.Epics = make([]core.Epic, 0)
	err = tx.Stmtx(s.stmt["selectProjectEpics"]).SelectContext(ctx, &p.Epics, p.ID)
	if err != nil && err != sql.ErrNoRows {
		log.Println("selectProject 2", err)
		tx.Rollback()
		return err
	}

	return tx.Commit()

}
