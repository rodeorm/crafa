package filler

import (
	"context"
	"money/internal/core"
)

type MessageStorager interface {
	UpdateMsg(ctx context.Context, m *core.Message) error
	SelectUnsendedMsgs(context.Context) ([]core.Message, error)
}
