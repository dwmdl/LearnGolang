package main

import (
	"api/configs"
	"api/internal/auth"
	"fmt"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.HandlerDeps{Config: config})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")

	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}
