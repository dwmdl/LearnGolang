package main

import (
	"api/configs"
	"api/internal/auth"
	"api/internal/link"
	"api/internal/stat"
	"api/internal/user"
	"api/pkg/db"
	"api/pkg/middleware"
	"fmt"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	router := http.NewServeMux()
	database := db.NewDb(config)

	// Repositories
	linkRepo := link.NewRepository(database)
	userRepo := user.NewRepository(database)
	statRepo := stat.NewRepository(database)

	//Services
	authService := auth.NewService(userRepo)

	// Handlers
	auth.NewHandler(router, auth.HandlerDeps{
		Config:  config,
		Service: authService,
	})
	link.NewHandler(router, link.HandlerDeps{
		LinkRepo: linkRepo,
		StatRepo: statRepo,
		Config:   config,
	})

	// Middlewares
	middlewareStack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: middlewareStack(router),
	}

	fmt.Println("Server is listening on port 8081")

	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}
