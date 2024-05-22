package item

type CreateItemDTO struct {
	Title  string
	Types  []int
	Prices []int
	Rating int
	Image  []string
	Sizes  []int
	Props  []bool
}
