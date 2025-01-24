package core

import (
	"context"
	"database/sql"

	"money/internal/crypt"

	"gopkg.in/gomail.v2"
)

const (
	MessageConfirm = iota + 1
	MessageAuth
	MessageNotify
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

	GMS *gomail.Message
}

type EmailStorager interface {
	AddEmail(ctx context.Context, m *Email) error
	SelectUnsendedEmails(context.Context) ([]Email, error)
	UpdateEmail(context.Context, *Email) error
}

// NewAuthMessage создает новое сообщение для аутентификации
func NewAuthEmail(u User) *Email {
	return &Email{User: u, Text: crypt.GetOneTimePassword()}
}

// NewConfirmEmail создает новое сообщение для подтверждения адреса электронной почты
func NewConfirmEmail(u User, url string) *Email {
	return &Email{User: u, Text: crypt.GetVerifyURL(url)}
}
