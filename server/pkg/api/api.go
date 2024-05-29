package api

import (
	"PizzaApi/config"
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
	cfg := config.New()
	client, err := postgres.NewClient(ctx, cfg.MaxAttempts, time.Duration(cfg.Delay)*time.Second, cfg.Username, cfg.Password, cfg.Database, cfg.Host, cfg.Port)
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
