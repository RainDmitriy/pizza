package order

import (
	"PizzaApi/internal/cartItem"
)

type Order struct {
	Id         int
	TotalPrice int
	CartItem   []cartItem.CartItem
}
