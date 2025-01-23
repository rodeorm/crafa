package cfg

import (
	"time"
)

type AppConfig struct {
	RunAddress      string        //Адрес запуска
	Domain          string        //Домен
	SSLPath         string        //Путь к сертификату SSL
	SSLKey          string        //Путь к ключу SSL
	ReadTimeout     time.Duration //Таймаут на чтение, сек
	WriteTimeout    time.Duration //Таймаут на запись, сек
	ShutdownTimeout time.Duration //Таймаут на выключение, сек
}
