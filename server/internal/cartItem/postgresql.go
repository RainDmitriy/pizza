package cartItem

import (
	"PizzaApi/pkg/client/postgres"
	"context"
	"log"
)

type repository struct {
	client postgres.Client
}

func (r *repository) GetAll(ctx context.Context) ([]CartItem, error) {
	q := `SELECT cart.cartId, cart.itemId, cart.selectedType, cart.selectedSize, cart.quantity,
							 items.title, 
							 Array[prices.forSmall, prices.forMiddle, prices.forBig] AS prices, 
							 Array[image.thin, image.traditional] AS image
				FROM cart
				JOIN items ON cart.itemId = items.itemId
				JOIN prices ON cart.itemId = prices.itemId
				JOIN image ON cart.itemId = image.itemId
				ORDER BY cartId;
			 `
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		log.Default().Println(err)
		return []CartItem{}, nil
	}
	items := make([]CartItem, 0)
	for rows.Next() {
		var item CartItem
		err = rows.Scan(&item.CartId, &item.ItemId, &item.SelectedType, &item.SelectedSize,
			&item.Quantity, &item.Title, &item.Prices, &item.Image)
		if err != nil {
			return []CartItem{}, err
		}
		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		log.Default().Println(err)
		return []CartItem{}, err
	}
	return items, nil
}

func (r *repository) GetOne(ctx context.Context, id int) (CartItem, error) {
	q := `SELECT cart.cartId, cart.itemId, cart.selectedType, cart.selectedSize, cart.quantity,
					items.title, 
					Array[prices.forSmall, prices.forMiddle, prices.forBig] AS prices, 
					Array[image.thin, image.traditional] AS image
				FROM cart
					JOIN items ON cart.itemId = items.itemId
					JOIN prices ON cart.itemId = prices.itemId
					JOIN image ON cart.itemId = image.itemId
				WHERE cart.CartId = $1;
			 `
	var item CartItem
	err := r.client.QueryRow(ctx, q, id).Scan(&item.CartId, &item.ItemId, &item.SelectedType, &item.SelectedSize,
		&item.Quantity, &item.Title, &item.Prices, &item.Image)
	if err != nil {
		return CartItem{}, err
	}
	return item, nil
}
func (r *repository) Update(ctx context.Context, cartItem CartItem) (int, error) {
	securityQ := `SELECT itemId FROM items WHERE title = $1;`
	err := r.client.QueryRow(ctx, securityQ, cartItem.Title).Scan(&cartItem.ItemId)
	if err != nil {
		log.Default().Println(err)
		return -1, err
	}
	q := `UPDATE cart SET cartId = $1, itemId = $2, selectedType = $3, selectedSize = $4, quantity = $5
				WHERE cartId = $1
				RETURNING cartId;`
	err = r.client.QueryRow(ctx, q, cartItem.CartId, cartItem.ItemId, cartItem.SelectedType, cartItem.SelectedSize, cartItem.Quantity).Scan(&cartItem.CartId)
	if err != nil {
		log.Default().Println(err)
		return -1, err
	}
	return cartItem.CartId, nil
}
func (r *repository) Delete(ctx context.Context, id int) (int, error) {
	q := `DELETE FROM cart
				WHERE cartId = $1
				RETURNING cartId;`
	err := r.client.QueryRow(ctx, q, id).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}
func (r *repository) Insert(ctx context.Context, cartItem CartItem) (CartItem, error) {
	q := `INSERT INTO cart 
				(cartId, itemId, selectedType, selectedSize, quantity)
				VALUES ($1, $2, $3, $4, $5)
				RETURNING cartId;`
	err := r.client.QueryRow(ctx, q, cartItem.CartId, cartItem.ItemId, cartItem.SelectedType, cartItem.SelectedSize, cartItem.Quantity).Scan(&cartItem.CartId)
	if err != nil {
		return CartItem{}, err
	}
	// echo
	return cartItem, nil
}

func NewRepository(client postgres.Client) Repository {
	return &repository{
		client: client,
	}
}
