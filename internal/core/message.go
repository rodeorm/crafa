package core

import (
	"context"
	"database/sql"
)

const (
	MessageTypeConfirm = iota + 1
	MessageTypeAuth
	MessageTypeNotify
)

const (
	MessageCategoryEmail = iota
	MessageCategorySMS
)

type MessageStorager interface {
	UpdateMsg(ctx context.Context, m *Message) error
	SelectUnsendedMsgs(context.Context) ([]Message, error)
}

// Базовое сообщение (может быть основой для email, sms, push и т.п.)
type Message struct {
	User
	Type
	MessageCategory

	SendTime  sql.NullTime // Время, когда сообщение было отправлено
	QueueTime sql.NullTime // Время, когда сообщение было помещенено в очередь
	Text      string       // Сообщение
	ID        int
	Used      bool // OTP из сообщения уже был использован
	Queued    bool // Сообщение в очереди на отправку
}

type Type struct {
	ID    int `db:"type.id"`
	Name  string
	Const string
}

type MessageCategory struct {
	ID    int `db:"category.id"`
	Name  string
	Const string
}
