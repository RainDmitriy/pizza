package cartItem

import (
	"context"
)

type Repository interface {
	GetAll(ctx context.Context) ([]CartItem, error)
	GetOne(ctx context.Context, id int) (CartItem, error)
	Update(ctx context.Context, cartItem CartItem) (int, error)
	Delete(ctx context.Context, id int) (int, error)
	Insert(ctx context.Context, cartItem CartItem) (CartItem, error)
}
