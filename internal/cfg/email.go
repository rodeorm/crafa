package cfg

import (
	"crafa/internal/queue"
)

type EmailConfig struct {
	//Количество наполнителей очереди на отправку
	FillWorkerCount int `yaml:"FILL_WORKERS"`
	//Количество отправителей
	SendWorkerCount int `yaml:"SEND_WORKERS"`
	//Адрес сервера электронной почты
	SMTPServer string `yaml:"SMTP_SERVER"`
	//Порт сервера электронной почты
	SMTPPort int `yaml:"SMTP_PORT"`
	//Логин сервера электронной почты
	SMTPLogin string `yaml:"SMTP_LOGIN"`
	//Пароль сервера электронной почты
	SMTPPass string `yaml:"SMTP_PASSWORD"`
	//Периодичность отправки сообщений (В секундах)
	MessageSendPeriod int `yaml:"MESSAGE_SEND_PERIOD"`
	//Периодичность наполнения очереди на отправку (В секундах)
	QueueFillPeriod int `yaml:"QUEUE_FILL_PERIOD"`
	//От кого
	From string `yaml:"FROM"`
	// Файл-вложение
	File string `yaml:"FILE"`

	Queue *queue.MessageQueue // Реализация очереди
}
