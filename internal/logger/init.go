package logger

import "go.uber.org/zap"

func init() {
	lvl, err := zap.ParseAtomicLevel("INFO")
	if err != nil {
		return
	}

	cfg := zap.NewProductionConfig()
	cfg.Level = lvl
	zl, err := cfg.Build()
	if err != nil {
		return
	}

	Log = zl
}
