// Package logger отражает работу с логгированием в проекте
package logger

import (
	"go.uber.org/zap"
)

// Log будет доступен всему коду как синглтон.
var Log *zap.Logger = zap.NewNop()
