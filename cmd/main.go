package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo/api"
	"todo/config"
	"todo/service"
	"todo/storage"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// Initialize storage (MongoDB connection)
	store, err := storage.NewStorage(cfg.DBUri) // Make sure NewStorage initializes your MongoDB client properly
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	// Initialize services
	userService := service.NewUserService(store.User())
	taskService := service.NewTaskService(store.Task())
	taskListService := service.NewTaskListService(store.TaskList())
	labelService := service.NewLabelService(store.Label())

	// Create a new Service struct as a pointer
	services := &service.Service{
		UserService:     userService,
		TaskService:     taskService,
		TaskListService: taskListService,
		LabelService:    labelService,
	}

	// Set up the router and server, passing services as a pointer
	r := api.New(services, log.Default())

	// Create a new server
	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	// Channel to listen for interrupt signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Run server in a goroutine
	go func() {
		fmt.Printf("Server is running on port %s\n", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("could not listen on %s: %v\n", cfg.Port, err)
		}
	}()

	// Wait for interrupt signal
	<-signalChan

	// Create a context with timeout to allow existing requests to finish
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown server gracefully
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
