package main

import (
	"api/configs"
	"api/internal/auth"
	"api/internal/link"
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

	// Handlers
	auth.NewAuthHandler(router, auth.HandlerDeps{Config: config})
	link.NewLinkHandler(router, link.HandlerDeps{LinkRepo: linkRepo})

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
