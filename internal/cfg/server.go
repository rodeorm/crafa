package cfg

import (
	"money/internal/logger"

	"go.uber.org/zap"
)

type ServerConfig struct {
	*AppConfig
	*SecurityConfig
}

func ConfigurateServer() (*ServerConfig, error) {
	a, _, _, s := Configurate()

	logger.Log.Info("Сконфигурировали http сервер",
		zap.String("Адрес запуска сервера", appCfg.RunAddress),
		zap.String("Пусть к ssl ключу", a.SSLKey),
		zap.String("Путь к ssl сертификату", a.SSLPath),
		zap.String("Таймаут на чтение", a.ReadTimeout.String()),
		zap.String("Таймаут на запись", a.WriteTimeout.String()),
	)

	return &ServerConfig{AppConfig: a, SecurityConfig: s}, nil
}
