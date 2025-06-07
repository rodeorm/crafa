package main

import (
	"crafa/internal/cfg"
	"crafa/internal/http/server"
	"crafa/internal/logger"
	"crafa/internal/msg/filler"
	"crafa/internal/msg/sender"
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
