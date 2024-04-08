package main

import (
	"log"

	"money/internal/auth"
	"money/internal/db"
	"money/internal/ui"
)

func main() {
	config, err := configuate()
	if err != nil {
		log.Println("main", err)
		return
	}

	wpDB, err := db.InitPostgres(config.WorkplaceDB)
	defer wpDB.CloseConnection()
	if err != nil {
		log.Println("main. WP", err)
		return
	}

	authDB, err := auth.InitPostgres(config.AuthDB)
	defer authDB.CloseConnection()
	if err != nil {
		log.Println("main. Auth", err)
		return
	}

	ui.StartUI(authDB, wpDB, config.RunAddress)
}
