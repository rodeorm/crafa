package cfg

import (
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var (
	appCfg      *AppConfig
	messageCfg  *MessageConfig
	postgresCfg *PostgresConfig
	securityCfg *SecurityConfig

	once sync.Once
)

func Configurate() (*AppConfig, *MessageConfig, *PostgresConfig, *SecurityConfig) {
	once.Do(
		func() {
			appCfg = &AppConfig{
				RunAddress:      os.Getenv("RUN_ADDRESS"),
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

			messageCfg = &MessageConfig{
				FillWorkerCount:   runtime.NumCPU() / 2,
				SendWorkerCount:   runtime.NumCPU() / 2,
				SMTPServer:        os.Getenv("SMTP_SERVER"),
				SMTPPort:          smtpPort,
				SMTPLogin:         os.Getenv("SMTP_LOGIN"),
				SMTPPass:          os.Getenv("SMTP_PASS"),
				MessageSendPeriod: 30,
				QueueFillPeriod:   30,
			}

			postgresCfg = &PostgresConfig{
				ConnectionString: os.Getenv("POSTGRES_CONNECTION"),
			}

			securityCfg = &SecurityConfig{
				TokeLiveTime: 1800000,
				JWTKey:       os.Getenv("JWK_KEY"),
			}

		})

	return appCfg, messageCfg, postgresCfg, securityCfg
}
