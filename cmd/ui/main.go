package main

import (
	"money/internal/cfg"
	"money/internal/http/server"
	"sync"
)

func main() {
	// Через этот канал основные горутины узнают, что надо закрываться для изящного завершения работы
	exit := make(chan struct{})

	c, err := cfg.ConfigurateServer()
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	server.Start(c, &wg, exit)
	wg.Wait()
}
