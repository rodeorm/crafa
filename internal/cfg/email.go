package cfg

import "money/internal/core"

type EmailConfig struct {
	FillWorkerCount   int    `yaml:"FILL_WORKERS"`
	SendWorkerCount   int    `yaml:"SEND_WORKERS"`
	SMTPServer        string `yaml:"SMTP_SERVER"`         //Адрес сервера электронной почты
	SMTPPort          int    `yaml:"SMTP_PORT"`           //Порт сервера электронной почты
	SMTPLogin         string `yaml:"SMTP_LOGIN"`          //Логин сервера электронной почты
	SMTPPass          string `yaml:"SMTP_PASSWORD"`       //Пароль сервера электронной почты
	MessageSendPeriod int    `yaml:"MESSAGE_SEND_PERIOD"` //Периодичность отправки сообщений (В секундах)
	QueueFillPeriod   int    `yaml:"QUEUE_FILL_PERIOD"`   //Периодичность наполнения очереди на отправку (В секундах)
	From              string `yaml:"FROM"`                //Периодичность отправки сообщений (В секундах)
	File              string `yaml:"FILE"`                //Периодичность наполнения очереди на отправку (В секундах)

	EmailQueue *core.Queue
}
