package cfg

import (
	"time"
)

type AppConfig struct {
	RunAddress      string        `yaml:"RUN_ADDRESS"`                       //Адрес запуска
	SSLPath         string        `yaml:"SSL_SERTIFICATE_RELATIVE_PATH"`     //Путь к сертификату SSL
	SSLKey          string        `yaml:"SSL_SERTIFICATE_KEY_RELATIVE_PATH"` //Путь к ключу SSL
	ReadTimeout     time.Duration `yaml:"READ_TIMEOUT"`                      //Таймаут на чтение
	WriteTimeout    time.Duration `yaml:"WRITE_TIMEOUT"`                     //Таймаут на запись
	ShutdownTimeout time.Duration `yaml:"SHUTDOWN_TIMEOUT"`                  //Таймаут на выключение
}
