package cfg

import (
	"crafa/internal/logger"
	"crafa/internal/queue"
	"os"
	"strconv"
	"sync"
	"time"
)

type Config struct {
	AppConfig
	PostgresConfig
	EmailConfig
	SecurityConfig
}

var (
	cfg  *Config
	exit chan struct{} // Через этот канал основные горутины узнают, что надо закрываться для изящного завершения работы
	wg   sync.WaitGroup
	once sync.Once
)

func Configurate() (*Config, chan struct{}, *sync.WaitGroup) {
	once.Do(
		func() {
			cfg = &Config{}

			cfg.AppConfig = AppConfig{
				RunAddress:      os.Getenv("RUN_ADDRESS"),
				Domain:          "crafa.ru", //os.Getenv("DOMAIN")
				SSLPath:         "ssl/certificate_ca.crt",
				SSLKey:          "ssl/certificate.key",
				ReadTimeout:     10 * time.Second,
				WriteTimeout:    10 * time.Second,
				ShutdownTimeout: 10 * time.Second,
			}

			smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
			if err != nil {
				smtpPort = 465
			}

			cfg.EmailConfig = EmailConfig{
				FillWorkerCount:   1, //runtime.NumCPU() / 2,
				SendWorkerCount:   3, //runtime.NumCPU() / 2,
				SMTPServer:        os.Getenv("SMTP_SERVER"),
				SMTPPort:          smtpPort,
				SMTPLogin:         os.Getenv("SMTP_LOGIN"),
				SMTPPass:          os.Getenv("SMTP_PASS"),
				MessageSendPeriod: 1,
				QueueFillPeriod:   1,
				Queue:             queue.NewQueue(5),
				From:              "i@ilyinal.ru",
				File:              "",
			}

			cfg.PostgresConfig = PostgresConfig{
				ConnectionString: os.Getenv("POSTGRES_CONNECTION"),
			}

			cfg.SecurityConfig = SecurityConfig{
				TokenLiveTime: 1800000, // os.Getenv("TOKEN_LIVE_TIME")
				OTPLiveTime:   18000,
				JWTKey:        os.Getenv("JWT_KEY"),
			}

			exit = make(chan struct{})

			logger.Sugar.Info("Запущена конфигурация. JWTKey", cfg.JWTKey)
		})

	return cfg, exit, &wg
}
