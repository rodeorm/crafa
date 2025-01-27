package main

import (
	"money/internal/cfg"
	"money/internal/http/server"
	"money/internal/msg/filler"
	"money/internal/msg/sender"
)

func main() {
	config, stgs, exit, wg := cfg.GetConfig()

	wg.Add(1 + config.EmailConfig.FillWorkerCount + config.EmailConfig.SendWorkerCount)
	go server.Start(config, stgs, wg, exit)
	go sender.Start(config, stgs.MessageStorager, wg, exit)
	go filler.Start(config, stgs.MessageStorager, wg, exit)
	wg.Wait()
}
