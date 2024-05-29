package cartItem

type CreateCartItemDTO struct {
	// quantity does not need, because it is 1 when created
	SelectedSize int
	SelectedType int
	ItemId       int
}
