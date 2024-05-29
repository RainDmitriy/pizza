package item

import (
	"PizzaApi/pkg/client/postgres"
	"context"
)

type repository struct {
	client postgres.Client
}

func (r *repository) GetAll(ctx context.Context) ([]Item, error) {
	q := `SELECT i.ItemId, i.Title, i.Rating, i.types,
					ARRAY[p.forSmall, p.forMiddle, p.forBig] AS Prices,
					ARRAY[ps.Spicy, ps.Meat, ps.Vegetarian, ps.Grilled, ps.Closed] AS Props,
					ARRAY[s.small, s.middle, s.big] AS Sizes,
					ARRAY[im.thin, im.traditional] AS Image
				FROM items i
				JOIN prices p ON i.ItemId = p.ItemId
				JOIN props ps ON i.ItemId = ps.ItemId
				JOIN sizes s ON i.ItemId = s.ItemId
				JOIN image im ON i.ItemId = im.ItemId
				ORDER BY i.ItemId;
			 `

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	items := make([]Item, 0)
	for rows.Next() {
		var item Item
		err = rows.Scan(&item.ItemId, &item.Title, &item.Rating, &item.Types, &item.Prices, &item.Props, &item.Sizes, &item.Image)
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

func (r *repository) GetOne(ctx context.Context, id int) (Item, error) {
	q := `SELECT i.ItemId, i.Title, i.Rating, i.types,
					ARRAY[p.forSmall, p.forMiddle, p.forBig] AS Prices,
					ARRAY[ps.Spicy, ps.Meat, ps.Vegetarian, ps.Grilled, ps.Closed] AS Props,
					ARRAY[s.small, s.middle, s.big] AS Sizes,
					ARRAY[im.thin, im.traditional] AS Image
				FROM items i
				JOIN prices p ON i.ItemId = p.ItemId
				JOIN props ps ON i.ItemId = ps.ItemId
				JOIN sizes s ON i.ItemId = s.ItemId
				JOIN image im ON i.ItemId = im.ItemId
				WHERE i.ItemId = $1;
			 `
	var item Item
	err := r.client.QueryRow(ctx, q, id).Scan(&item.ItemId, &item.Title, &item.Rating, &item.Types, &item.Prices, &item.Props, &item.Sizes, &item.Image)
	if err != nil {
		return Item{}, err
	}
	return item, nil
}

func NewRepository(client postgres.Client) Repository {
	return &repository{
		client: client,
	}
}
