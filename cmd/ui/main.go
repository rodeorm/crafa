package main

import (
	"money/internal/cfg"
	"money/internal/http/server"
	"sync"
)

func main() {
	// Через этот канал горутины узнают, что надо закрываться для изящного завершения работы
	exit := make(chan struct{})

	s := cfg.ServerConfig{
		AppConfig: cfg.AppConfig{
			RunAddress:      "localhost:8080",
			ReadTimeout:     10,
			WriteTimeout:    10,
			ShutdownTimeout: 10,
		},
	}

	var wg sync.WaitGroup

	server.Start(s, &wg, exit)

}
