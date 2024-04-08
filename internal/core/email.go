package core

import (
	"database/sql"
)

type Email struct {
	Login            string       // Логин, для которого было отправлено подтверждение
	Email            string       // Адрес электронной почты, на который было отправлено подтверждение
	Key              string       // Код для ссылки, подтверждающий email
	SendedTime       sql.NullTime // Время, когда сообщение было отправлено
	ConfirmedTime    sql.NullTime // Время, когда email был подтержден
	ConfirmationLink string       // Ссылка, использованная для подтверждения
	TemplateID       int

	User
}

type EmailTemplate struct {
	ID   int
	Name string
	Text string
}
