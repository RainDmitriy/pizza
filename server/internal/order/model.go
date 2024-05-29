package order

import (
	"PizzaApi/internal/cartItem"
)

type Order struct {
	OrderId    int
	TotalPrice int
	CartItems  []cartItem.CartItem
}
