package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Product struct {
	ProductPK int64                         `json:"id"`
	Data      map[string]ProductInformation `json:"data"`
	UpdatedAt time.Time                     `json:"-"`
}

type ProductModel struct {
	DB *sql.DB
}

func (m ProductModel) Insert(product *Product) error {
	query := `
        INSERT INTO products (product_pk, data, updated_at) 
        VALUES ($1, $2, $3)
        RETURNING product_pk, data, updated_at`

	args := []any{product.ProductPK, product.Data, product.UpdatedAt}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&product.ProductPK, &product.Data, &product.UpdatedAt)
}

func (m ProductModel) Get(id int64) (*Product, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
        SELECT product_pk, data, updated_at
        FROM products
        WHERE product_pk = $1`

	var product Product

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&product.ProductPK,
		&product.Data,
		&product.UpdatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &product, nil
}

func (m ProductModel) Update(product *Product) error {
	query := `
        UPDATE products 
        SET data = $1, updated_at = $2
        WHERE product_pk = $3
        RETURNING product_pk`

	args := []any{
		product.ProductPK,
		product.Data,
		product.UpdatedAt,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&product.ProductPK)
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

func (m ProductModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
        DELETE FROM products
        WHERE product_pk = $1`

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

func (m ProductModel) GetAll(productPk string, filters Filters) ([]*Product, Metadata, error) {
	query := fmt.Sprintf(`
        SELECT count(*) OVER(), product_pk, data, updated_at
        FROM products
        WHERE (to_tsvector('simple', product_pk) @@ plainto_tsquery('simple', $1) OR $1 = '') 
        ORDER BY %s %s, product_pk ASC
        LIMIT $2 OFFSET $3`, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{productPk, filters.limit(), filters.offset()}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	pArr := []*Product{}

	for rows.Next() {
		var p Product

		err := rows.Scan(
			&totalRecords,
			&p.ProductPK,
			&p.Data,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		pArr = append(pArr, &p)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return pArr, metadata, nil
}
