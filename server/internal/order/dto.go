package order

import (
	"PizzaApi/internal/cartItem"
)

type CreateOrderDTO struct {
	TotalPrice int
	CartItem   []cartItem.CartItem
}
