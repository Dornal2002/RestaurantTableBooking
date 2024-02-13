package main

import (
	"context"
	"log"
	"net/http"
	"project/internal/api"
	"project/internal/app"
	"project/internal/repository"
)

func main() {
	ctx := context.Background()
	log.Print(ctx, "Starting Application")

	database, err := repository.InitializeDatabse()
	if err != nil {
		log.Fatal("Error Occured while Initializing database", err)
		return
	}
	services := app.NewServices(database)

	router := api.NewRouter(services)

	http.ListenAndServe("localhost:8080", router)
}
