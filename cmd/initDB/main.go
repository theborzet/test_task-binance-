package main

// это я добавил для удобства проверки, чтобы была понятна структура бд и не пришлось вручную ее создавать

import (
	"log"
	"test_api/internal/config"
	"test_api/internal/db"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Println("Some problems with config", err)
	}
	database := db.Init(config)
	defer func() {
		if err := db.Close(database); err != nil {
			log.Fatalf("Failed to close DB: %v", err)
		}
	}()
	db.MigrateDB(database)
}
