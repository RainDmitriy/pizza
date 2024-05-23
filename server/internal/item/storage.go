package item

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, item *Item) error
	FindAll(ctx context.Context) ([]Item, error)
	FindOne(ctx context.Context, id int) (Item, error)
	Update(ctx context.Context, id int, item *Item) (Item, error)
	Delete(ctx context.Context, id int) (Item, error)
}
