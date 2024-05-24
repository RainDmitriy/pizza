package item

import (
	"PizzaApi/pkg/client/postgres"
	"context"
	"fmt"

	"github.com/jackc/pgconn"
)

type repository struct {
	client postgres.Client
}

func (r *repository) Create(ctx context.Context, item *Item) error {
	q := `
				INSERT INTO items
					(title, types, price, rating, image, sizes, props)
				VALUES 
						($1, $2, $3, $4, $5, $6, $7) 
				RETURNING id
				`

	if err := r.client.QueryRow(ctx, q, item.Title, item.Types, item.Prices, item.Rating, item.Image, item.Sizes, item.Props).Scan(&item.Id); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf(fmt.Sprintf("pq: code: %s, message: %s, details: %s", pgErr.Code, pgErr.Message, pgErr.Detail))
			fmt.Println(newErr)
			return newErr
		}
		return err
	}
	return nil
}

func (r *repository) FindAll(ctx context.Context) ([]Item, error) {
	q := `SELECT 
					id, title, types, price, rating, image, sizes, props
				FROM items
				ORDER BY id
			 `

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	items := make([]Item, 0)
	for rows.Next() {
		var item Item
		err = rows.Scan(&item.Id, &item.Title, &item.Types, &item.Prices, &item.Rating, &item.Image, &item.Sizes, &item.Props)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *repository) FindOne(ctx context.Context, id int) (Item, error) {

	q := `SELECT 
					id, title, types, price, rating, image, sizes, props
				FROM items
				WHERE id = $1
				ORDER BY id
			 `
	var i Item
	err := r.client.QueryRow(ctx, q, id).Scan(&i.Id, &i.Title, &i.Types, &i.Prices, &i.Rating, &i.Image, &i.Sizes, &i.Props)

	if err != nil {
		return Item{}, err
	}
	return i, nil
}
func (r *repository) Update(ctx context.Context, id int, updatedItem *Item) (Item, error) {
	q := `UPDATE items
			SET title = $1, types = $2, price = $3, rating = $4, image = $5, sizes = $6, props = $7
			WHERE id = $8
			RETURNING id, title, types, price, rating, image, sizes, props
			`
	var i Item
	err := r.client.QueryRow(ctx, q, updatedItem.Title, updatedItem.Types, updatedItem.Prices, updatedItem.Rating, updatedItem.Image, updatedItem.Sizes, updatedItem.Props, id).Scan(&i.Id, &i.Title, &i.Types, &i.Prices, &i.Rating, &i.Image, &i.Sizes, &i.Props)

	if err != nil {
		return Item{}, err
	}
	return i, nil
}

func (r *repository) Delete(ctx context.Context, id int) (Item, error) {
	q := `DELETE FROM items
			WHERE id = $1
			RETURNING id, title, types, price, rating, image, sizes, props
			`
	var i Item
	err := r.client.QueryRow(ctx, q, id).Scan(&i.Id, &i.Title, &i.Types, &i.Prices, &i.Rating, &i.Image, &i.Sizes, &i.Props)

	if err != nil {
		return Item{}, err
	}
	return i, nil
}

func NewRepository(client postgres.Client) Repository {
	return &repository{
		client: client,
	}
}
