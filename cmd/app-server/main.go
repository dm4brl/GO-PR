package main

import (
	"go-automation/internal/api"
	"go-automation/pkg/config"
	"log"
)

func main() {
	cfg := config.Load()

	// Запускаем API сервер
	if err := api.StartHTTPServer(cfg); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}
