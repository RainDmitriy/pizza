package api

import (
	"PizzaApi/internal/cartItem"
	"PizzaApi/internal/item"
	"PizzaApi/internal/order"
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
	// as well we need to move this to config.yaml and use in config.go
	delay := 5 * time.Second
	maxAttempts := 5
	login := "itmo409189_2024"
	password := "itmo409189"
	database := "dbstud"
	host := "146.185.211.205"
	port := "5432"
	client, err := postgres.NewClient(ctx, maxAttempts, delay, login, password, database, host, port)
	if err != nil {
		log.Fatal(err)
	}
	itemHandler := item.Handler(client)
	cartHandler := cartItem.Handler(client)
	orderHadler := order.Handler(client)

	api.r.HandleFunc("/item/{id}", itemHandler)
	api.r.HandleFunc("/cart/{id}", cartHandler)
	api.r.HandleFunc("/orders/{id}", orderHadler)

}

func (api *api) ListenAndServe() error {

	return http.ListenAndServe(api.addr, api.r)

}
