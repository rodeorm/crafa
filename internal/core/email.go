package core

import (
	"bytes"
	"fmt"
	"path/filepath"
	"text/template"

	"gopkg.in/gomail.v2"
)

// Email сообщение
type Email struct {
	Message
	GMS *gomail.Message
}

// NewEmail создает новое письмо с набором функц опций
func NewEmail(m Message, opts ...func(*Email)) *Email {
	e := &Email{GMS: gomail.NewMessage(),
		Message: m}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

// WithAttachment добавляет к письму вложение
func WithAttachment(filePath string) func(*Email) {
	return func(e *Email) {
		attachPath, err := filepath.Abs(filepath.Join(".", "static", "img", filePath))
		if err == nil {
			return
		}
		e.GMS.Attach(attachPath)
	}
}

// WithHeader добавляет subject, from, to
func WithHeader(from, to string) func(*Email) {
	return func(e *Email) {
		switch e.Type.ID {
		case MessageTypeConfirm:
			e.GMS.SetHeader("Subject", "Подтверждение адреса электронной почты")
		case MessageTypeAuth:
			e.GMS.SetHeader("Subject", "Одноразовый пароль")
		case MessageTypeNotify:
			e.GMS.SetHeader("Subject", "Уведомление")
		}
		e.GMS.SetHeader("From", from)
		e.GMS.SetHeader("To", to)
	}
}

// WithBody персонализирует текст email сообщения по шаблону (папка, страница)
func WithBody(domain, otp string) func(*Email) {
	return func(e *Email) {
		var (
			templatePath string
			body         bytes.Buffer
		)
		switch e.Type.ID {
		case MessageTypeConfirm:
			templatePath, _ = filepath.Abs(fmt.Sprintf("./view/%s/%s.html", "email", "confirm"))
		case MessageTypeAuth:
			templatePath, _ = filepath.Abs(fmt.Sprintf("./view/%s/%s.html", "email", "auth"))
		case MessageTypeNotify:
			templatePath, _ = filepath.Abs(fmt.Sprintf("./view/%s/%s.html", "email", "notify"))
		}

		mail, _ := template.ParseFiles(templatePath)
		mail.Execute(&body, otp)
		e.GMS.SetBody("text/html", body.String())
	}
}
