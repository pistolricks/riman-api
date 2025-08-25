package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/pistolricks/riman-api/internal/validator"

	"github.com/lib/pq"
)

type Order struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitzero"`
	Runtime   Runtime   `json:"runtime,omitzero"`
	Genres    []string  `json:"genres,omitzero"`
	Version   int32     `json:"version"`
}

func ValidateOrder(v *validator.Validator, order *Order) {
	v.Check(order.Title != "", "title", "must be provided")
	v.Check(len(order.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(order.Year != 0, "year", "must be provided")
	v.Check(order.Year >= 1888, "year", "must be greater than 1888")
	v.Check(order.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(order.Runtime != 0, "runtime", "must be provided")
	v.Check(order.Runtime > 0, "runtime", "must be a positive integer")

	v.Check(order.Genres != nil, "genres", "must be provided")
	v.Check(len(order.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(order.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(order.Genres), "genres", "must not contain duplicate values")
}

type OrderModel struct {
	DB *sql.DB
}

func (m OrderModel) Insert(order *Order) error {
	query := `
        INSERT INTO orders (title, year, runtime, genres) 
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at, version`

	args := []any{order.Title, order.Year, order.Runtime, pq.Array(order.Genres)}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&order.ID, &order.CreatedAt, &order.Version)
}

func (m OrderModel) Get(id int64) (*Order, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
        SELECT id, created_at, title, year, runtime, genres, version
        FROM orders
        WHERE id = $1`

	var order Order

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&order.ID,
		&order.CreatedAt,
		&order.Title,
		&order.Year,
		&order.Runtime,
		pq.Array(&order.Genres),
		&order.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &order, nil
}

func (m OrderModel) Update(order *Order) error {
	query := `
        UPDATE orders 
        SET title = $1, year = $2, runtime = $3, genres = $4, version = version + 1
        WHERE id = $5 AND version = $6
        RETURNING version`

	args := []any{
		order.Title,
		order.Year,
		order.Runtime,
		pq.Array(order.Genres),
		order.ID,
		order.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&order.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

func (m OrderModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
        DELETE FROM orders
        WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (m OrderModel) GetAll(title string, genres []string, filters Filters) ([]*Order, Metadata, error) {
	query := fmt.Sprintf(`
        SELECT count(*) OVER(), id, created_at, title, year, runtime, genres, version
        FROM orders
        WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '') 
        AND (genres @> $2 OR $2 = '{}')     
        ORDER BY %s %s, id ASC
        LIMIT $3 OFFSET $4`, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{title, pq.Array(genres), filters.limit(), filters.offset()}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	orders := []*Order{}

	for rows.Next() {
		var order Order

		err := rows.Scan(
			&totalRecords,
			&order.ID,
			&order.CreatedAt,
			&order.Title,
			&order.Year,
			&order.Runtime,
			pq.Array(&order.Genres),
			&order.Version,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		orders = append(orders, &order)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return orders, metadata, nil
}
