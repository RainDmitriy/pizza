package cartItem

type CartItem struct {
	CartId       int
	ItemId       int
	Quantity     int
	SelectedSize int
	SelectedType int
	Title        string
	Prices       []int
	Image        []string
}
