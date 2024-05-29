package item

import (
	"context"
)

type Repository interface {
	GetAll(ctx context.Context) ([]Item, error)
	GetOne(ctx context.Context, id int) (Item, error)
}
