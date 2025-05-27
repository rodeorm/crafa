package main

import (
	"money/internal/cfg"
	"money/internal/http/server"
	"money/internal/logger"
	"money/internal/msg/filler"
	"money/internal/msg/sender"
)

func main() {
	config, exit, wg := cfg.Configurate()

	defer logger.Log.Sync()
	wg.Add(1 + config.EmailConfig.FillWorkerCount + config.EmailConfig.SendWorkerCount)

	go server.Start(config, wg, exit)
	go sender.Start(config, wg, exit)
	go filler.Start(config, wg, exit)

	wg.Wait()
}
