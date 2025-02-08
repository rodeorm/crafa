package logger

import (
	"sync"

	"go.uber.org/zap"
)

// LoggerWrapper обертка над zap.Logger с синхронизацией
type LoggerWrapper struct {
	logger *zap.Logger
	mu     sync.Mutex
}

// NewLoggerWrapper создает новый экземпляр LoggerWrapper
func NewLoggerWrapper() (*LoggerWrapper, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return &LoggerWrapper{logger: logger}, nil
}

// Info записывает информационное сообщение
func (lw *LoggerWrapper) Info(msg string, fields ...zap.Field) {
	lw.mu.Lock()
	defer lw.mu.Unlock()
	lw.logger.Info(msg, fields...)
}

// Error записывает сообщение об ошибке
func (lw *LoggerWrapper) Error(msg string, fields ...zap.Field) {
	lw.mu.Lock()
	defer lw.mu.Unlock()
	lw.logger.Error(msg, fields...)
}

// Warn записывает предупреждающее сообщение
func (lw *LoggerWrapper) Warn(msg string, fields ...zap.Field) {
	lw.mu.Lock()
	defer lw.mu.Unlock()
	lw.logger.Warn(msg, fields...)
}

// Fatal записывает фатальное сообщение и завершает приложение
func (lw *LoggerWrapper) Fatal(msg string, fields ...zap.Field) {
	lw.mu.Lock()
	defer lw.mu.Unlock()
	lw.logger.Fatal(msg, fields...)
}
