package sender

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"

	"money/internal/core"

	"gopkg.in/gomail.v2"
)

func (s Sender) NewEmail(from string, m *core.Email) *core.Email {
	text := customiseMail("mail", "otp", m)
	ems := &core.Email{GMS: gomail.NewMessage()}
	ems.GMS.SetHeader("From", s.from)
	ems.GMS.SetHeader("To", m.Email)
	ems.GMS.SetHeader("Subject", "Одноразовый пароль от keeper")
	ems.GMS.SetBody("text/html", text)
	attachPath, err := filepath.Abs(filepath.Join(".", "static", "img", s.fileName))
	if err == nil {
		ems.GMS.Attach(attachPath)
	}
	return ems
}

func customiseMail(folder string, page string, param interface{}) string {
	templatePath, _ := filepath.Abs(fmt.Sprintf("./%s/%s.html", folder, page))
	mail, _ := template.ParseFiles(templatePath)
	var body bytes.Buffer
	mail.Execute(&body, param)
	return body.String()
}
