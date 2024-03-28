package main

import (
	// "context"
	"log"
	"net/http"
	"project/internal/api"
	"project/internal/app"
	"project/internal/repository"

	"github.com/rs/cors"
)

func main() {
	// ctx := context.Background()
	database, err := repository.InitializeDatabse()
	if err != nil {
		log.Fatal("Error Occured while Initializing database", err)
		return
	}

	services := app.NewServices(database)
	router := api.NewRouter(services)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"*"},
	})

	log.Println("Server is running on port 8080 ...")
	err = http.ListenAndServe("localhost:8080", c.Handler(router))
	if err != nil {
		log.Fatal("Error Occured while Initializing database", err)
		return
	}
}
