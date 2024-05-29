package order

import (
	"PizzaApi/pkg/client/postgres"
	"context"
	"encoding/json"
)

type repository struct {
	client postgres.Client
}

func (r *repository) GetAll(ctx context.Context) ([]Order, error) {
	q := `SELECT OrderId, CartItems, TotalPrice from orders
				ORDER BY OrderId;
			 `
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	orders := make([]Order, 0)
	for rows.Next() {
		var order Order
		err = rows.Scan(&order.OrderId, &order.CartItems, &order.TotalPrice)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *repository) GetOne(ctx context.Context, id int) (Order, error) {
	q := `SELECT OrderId, CartItems, TotalPrice from orders
				WHERE OrderId = $1;
			 `
	var order Order
	err := r.client.QueryRow(ctx, q, id).Scan(&order.OrderId, &order.CartItems, &order.TotalPrice)
	if err != nil {
		return Order{}, err
	}
	return order, nil
}

func (r *repository) Insert(ctx context.Context, order Order) (Order, error) {
	q := `INSERT INTO orders (CartItems, TotalPrice) VALUES ($1, $2) RETURNING OrderId;`
	cartItems, err := json.Marshal(order.CartItems)
	if err != nil {
		return Order{}, err
	}
	err = r.client.QueryRow(ctx, q, cartItems, order.TotalPrice).Scan(&order.OrderId)
	if err != nil {
		return Order{}, err
	}
	return order, nil
}

func NewRepository(client postgres.Client) Repository {
	return &repository{
		client: client,
	}
}
