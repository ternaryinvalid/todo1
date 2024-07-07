package main

import (
	http_adapter "github.com/ternaryinvalid/todo1/internal/app/adapters/primary/http-adapter"
	os_signal_adapter "github.com/ternaryinvalid/todo1/internal/app/adapters/primary/os-signal-adapter"
	todo_repository "github.com/ternaryinvalid/todo1/internal/app/adapters/secondary/repositories/todo-repository"
	api_service "github.com/ternaryinvalid/todo1/internal/app/application/api-service"
	"github.com/ternaryinvalid/todo1/internal/pkg/config"
	"log"
)

func main() {
	cfg := config.New()

	// Initialize Repositories
	todoRepository := todo_repository.New(cfg.Adapters.Secondary.Databases.Todo)

	// Initialize Services
	apiService := api_service.New(todoRepository, todoRepository)

	httpAdapter := http_adapter.New(cfg.Adapters.Primary.HttpAdapter, apiService)

	go httpAdapter.Start()

	osSignalAdapter := os_signal_adapter.New()

	go osSignalAdapter.Start()

	// Graceful shutdown
	select {
	case err := <-httpAdapter.Notify():
		log.Println(err.Error(), "main")
	case err := <-osSignalAdapter.Notify():
		log.Println(err.Error(), "main")
	}
}
