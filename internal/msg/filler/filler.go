package filler

import (
	"money/internal/core"
)

// NewFiller создает новый Filler
// Каждый Filler может наполнять очередь
func NewFiller(queue *core.Queue, storage core.EmailStorager, prd int) *Filler {
	return &Filler{
		queue:         queue,
		emailStorager: storage,
		period:        prd,
	}
}
