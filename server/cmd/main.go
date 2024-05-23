package main

import (
	item "PizzaApi/internal/item/db"
	"PizzaApi/pkg/api"
	"PizzaApi/pkg/client/postgres"
	"context"
	"log"
	"net/http"
	"time"
)

func main() {
	api := api.New("localhost:8080", http.NewServeMux())
	api.FillEndpoints()

	ctx := context.Background()
	maxAttempts := 5
	delay := 5 * time.Second
	username := "postgres"
	password := "kali"
	database := "postgres"
	host := "localhost"
	port := "5432"

	postgreSQLClient, err := postgres.NewClient(ctx, maxAttempts, delay, username, password, database, host, port)

	if err != nil {
		log.Fatal(err)
	}

	itemRepository := item.NewRepository(postgreSQLClient)

	all, err := itemRepository.FindAll(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println(all)

	log.Fatal(api.ListenAndServe())
}
