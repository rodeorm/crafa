package core

import (
	"context"
	"database/sql"
)

type Email struct {
	Login            string // Логин, для которого было отправлено подтверждение
	Email            string // Адрес электронной почты, на который было отправлено подтверждение
	Key              string // Код для ссылки, подтверждающий email
	ConfirmationLink string // Ссылка, использованная для подтверждения
	TemplateID       int

	SendedTime    sql.NullTime // Время, когда сообщение было отправлено
	ConfirmedTime sql.NullTime // Время, когда email был подтержден

	User
}

type EmailTemplate struct {
	ID   int
	Name string
	Text string
}

type EmailStorager interface {
	AddEmail(context.Context, *Email) error
	UpdateEmail(context.Context, *Email) error
}
