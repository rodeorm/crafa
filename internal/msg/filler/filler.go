package filler

import (
	"crafa/internal/core"
)

// Filler - рабочий, заполняющий очередь сообщений
type Filler struct {
	msgStorager MessageStorager // Хранилище сообщений
	queue       QueueStorager   // Очередь сообщений
	ID          int             // Идентификатор воркера
	period      int             // Периодичность наполнения сообщений
}

type QueueStorager interface {
	Push(m *core.Message)
}

// NewFiller создает новый Filler
// Каждый Filler может наполнять очередь
func NewFiller(queue QueueStorager, storage MessageStorager, prd int) *Filler {
	return &Filler{
		queue:       queue,
		msgStorager: storage,
		period:      prd,
	}
}
