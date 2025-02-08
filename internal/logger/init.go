package logger

import "go.uber.org/zap"

var Log *zap.Logger = zap.NewNop() // Будет доступен всему коду как синглтон

func init() {
	lvl, err := zap.ParseAtomicLevel("INFO")
	if err != nil {
	}

	cfg := zap.NewProductionConfig()
	cfg.Level = lvl
	zl, err := cfg.Build()

	Log = zl
}
