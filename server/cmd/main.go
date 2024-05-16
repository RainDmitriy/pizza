package main

import (
	"PizzaApi/pkg/api"
	"log"
	"net/http"
)

func main() {
	api := api.New("localhost:8080", http.NewServeMux())
	api.FillEndpoints()
	log.Fatal(api.ListenAndServe())
}
