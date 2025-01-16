// Package logger отражает работу с логгированием в проекте
package logger

import (
	"go.uber.org/zap"
)

// Log будет доступен всему коду как синглтон.
// Никакой код, кроме функции Initialize, не должен модифицировать эту переменную.
var Log *zap.Logger = zap.NewNop()
