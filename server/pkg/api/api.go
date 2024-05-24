package api

import (
	"PizzaApi/internal/item"
	"PizzaApi/pkg/client/postgres"
	"context"
	"log"
	"net/http"
	"time"
)

type api struct {
	addr string
	r    *http.ServeMux
}

func New(addr string, r *http.ServeMux) *api {
	return &api{
		addr: addr,
		r:    r,
	}
}

func (api *api) FillEndpoints() {
	ctx := context.Background()
	delay := 5 * time.Second
	maxAttempts := 5
	login := "postgres"
	password := "kali"
	database := "postgres"
	host := "localhost"
	port := "5432"
	client, err := postgres.NewClient(ctx, maxAttempts, delay, login, password, database, host, port)
	if err != nil {
		log.Fatal(err)
	}
	itemHandler := item.GetItemHandler(client)

	api.r.HandleFunc("/item", itemHandler)

}

func (api *api) ListenAndServe() error {

	return http.ListenAndServe(api.addr, api.r)

}
