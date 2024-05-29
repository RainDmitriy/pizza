package order

import (
	"context"
)

type Repository interface {
	GetAll(ctx context.Context) ([]Order, error)
	GetOne(ctx context.Context, id int) (Order, error)
	Insert(ctx context.Context, order Order) (Order, error)
}
