package order

import (
	"PizzaApi/internal/cartItem"
)

type CreateOrderDTO struct {
	TotalPrice int
	CartItems  []cartItem.CartItem
}
