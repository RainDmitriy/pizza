package cartItem

type CreateCartItemDTO struct {
	// quantity does not need, because it is 1 when created
	selectedSize int
	selectedType int
	ItemId       int
}
