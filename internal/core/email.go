package core

import (
	"context"
	"database/sql"

	"money/internal/crypt"
)

const (
	Verify = iota
	OTP
	Notify
)

// Email базовое сообщение
// В text пишется персонализированная информация: OneTimePassword, URL для подтверждения адреса электронной почты, уведомления
type Email struct {
	User

	SendedDate sql.NullTime //Время, когда сообщение было отправлено
	Attachment string       //Путь к вложению
	Text       string       //Сообщение

	ID   int //Идентификатор
	Type int //Тип: Verify, OTP, Notify

	Used   bool // OTP из сообщения уже был использован
	Queued bool // Сообщение в очереди на отправку
}

type EmailStorager interface {
	AddEmail(ctx context.Context, m *Email) error
	SelectUnsendedEmails(context.Context) ([]Email, error)
	UpdateEmail(context.Context, *Email) error
}

// NewAuthMessage создает новое сообщение
func NewAuthEmail(u User) *Email {
	return &Email{User: u, Text: crypt.GetOneTimePassword()}
}

// NewVerifyEmail создает новое сообщение
func NewVerifyEmail(u User, url string) *Email {
	return &Email{User: u, Text: crypt.GetVerifyURL(url)}
}
