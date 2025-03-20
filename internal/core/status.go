package core

import (
	"context"
	"fmt"
)

const (
	StatusNew      = iota // Статус "Новый"
	StatusPlanned         // Статус "Запланирован"
	StatusRejected        // Статус "Отменен"
	StatusDelayed         // Статус "Отложен"
	StatusOngoing         // Статус "В работе"
	StatusDone            // Статус "Выполнен"
)

type Status struct {
	Name string
	ID   int
}

type StatusStorager interface {
	SelectStatus(context.Context, *Status) error
	SelectAllStatuses(context.Context) ([]Status, error)
}

type StatusCash struct {
}

func (sc *StatusCash) SelectAllStatuses(context.Context) ([]Status, error) {
	ls := make([]Status, 6)
	ls[0] = Status{ID: StatusNew, Name: "New"}
	ls[1] = Status{ID: StatusPlanned, Name: "Planned"}
	ls[2] = Status{ID: StatusRejected, Name: "Rejected"}
	ls[3] = Status{ID: StatusDelayed, Name: "Delayed"}
	ls[4] = Status{ID: StatusOngoing, Name: "Ongoing"}
	ls[5] = Status{ID: StatusDone, Name: "Done"}

	return ls, nil
}

func (sc *StatusCash) SelectStatus(ctx context.Context, s *Status) error {
	switch s.ID {
	case StatusNew:
		s.Name = "New"
	case StatusPlanned:
		s.Name = "Planned"
	case StatusRejected:
		s.Name = "Rejected"
	case StatusDelayed:
		s.Name = "Delayed"
	case StatusOngoing:
		s.Name = "Ongoing"
	case StatusDone:
		s.Name = "Done"
	default:
		return fmt.Errorf("некорректный статус")
	}

	return nil
}
