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

	server := http.Server{
		Addr:    ":8081",
		Handler: middleware.Logging(router),
	}

	fmt.Println("Server is listening on port 8081")

	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}
