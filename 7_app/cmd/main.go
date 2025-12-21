package main

import (
	"api/configs"
	"api/internal/auth"
	"api/internal/link"
	"api/internal/stat"
	"api/internal/user"
	"api/pkg/db"
	"api/pkg/event"
	"api/pkg/middleware"
	"fmt"
	"net/http"
)

func App() http.Handler {
	config := configs.LoadConfig()
	router := http.NewServeMux()
	database := db.NewDb(config)
	eventBus := event.NewEventBus()

	// Repositories
	linkRepo := link.NewRepository(database)
	userRepo := user.NewRepository(database)
	statRepo := stat.NewRepository(database)

	//Services
	authService := auth.NewService(userRepo)
	eventService := stat.NewService(&stat.ServiceDeps{
		EventBus:   eventBus,
		Repository: statRepo,
	})

	// Handlers
	auth.NewHandler(router, auth.HandlerDeps{
		Config:  config,
		Service: authService,
	})
	link.NewHandler(router, link.HandlerDeps{
		LinkRepo: linkRepo,
		EventBus: eventBus,
		Config:   config,
	})
	stat.NewHandler(router, stat.HandlerDeps{
		StatRepo: statRepo,
		Config:   config,
	})

	go eventService.AddDirection()

	// Middlewares
	middlewareStack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	return middlewareStack(router)
}

func main() {
	app := App()
	server := http.Server{
		Addr:    ":8081",
		Handler: app,
	}

	fmt.Println("Server is listening on port 8081")
	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}
