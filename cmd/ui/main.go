package main

import (
	"money/internal/cfg"
	"money/internal/http/server"
)

func main() {
	config, exit, wg := cfg.GetConfig()

	wg.Add(1)
	go server.Start(config, wg, exit)
	// go sender.Start(config, wg, exit)
	// go filler.Start(config, wg, exit)
	wg.Wait()
}
