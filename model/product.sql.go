// Code generated by sqlc. DO NOT EDIT.
// source: product.sql

package model

import (
	"context"
	"database/sql"
)

const checkDuplicateProductUnit = `-- name: CheckDuplicateProductUnit :many
SELECT product_unit from scm_product where product_unit=?
`

func (q *Queries) CheckDuplicateProductUnit(ctx context.Context, productUnit *string) ([]*string, error) {
	rows, err := q.db.QueryContext(ctx, checkDuplicateProductUnit, productUnit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*string
	for rows.Next() {
		var product_unit *string
		if err := rows.Scan(&product_unit); err != nil {
			return nil, err
		}
		items = append(items, product_unit)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const createProduct = `-- name: CreateProduct :execresult
INSERT INTO scm_product (
    product_unit, manufacturer_id, product_name, reg_user_id
) VALUES ( ?, ?, ?, ?)
`

type CreateProductParams struct {
	ProductUnit    *string `json:"productUnit"`
	ManufacturerID *int    `json:"manufacturerID"`
	ProductName    *string `json:"productName"`
	RegUserID      *int    `json:"regUserID"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createProduct,
		arg.ProductUnit,
		arg.ManufacturerID,
		arg.ProductName,
		arg.RegUserID,
	)
}

const getProductFromProductID = `-- name: GetProductFromProductID :one
SELECT product_id, product_unit, manufacturer_id, product_name, reg_datetime, reg_user_id, mod_datetime, mod_user_id FROM scm_product
WHERE product_id = ? LIMIT 1
`

func (q *Queries) GetProductFromProductID(ctx context.Context, productID int32) (ScmProduct, error) {
	row := q.db.QueryRowContext(ctx, getProductFromProductID, productID)
	var i ScmProduct
	err := row.Scan(
		&i.ProductID,
		&i.ProductUnit,
		&i.ManufacturerID,
		&i.ProductName,
		&i.RegDatetime,
		&i.RegUserID,
		&i.ModDatetime,
		&i.ModUserID,
	)
	return i, err
}

const getProductFromProductUnit = `-- name: GetProductFromProductUnit :one
SELECT (product_id, product_unit, product_name, manufacturer_id, reg_datetime, reg_user_id, mod_datetime, mod_user_id)
from scm_product where product_unit=?
`

func (q *Queries) GetProductFromProductUnit(ctx context.Context, productUnit *string) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, getProductFromProductUnit, productUnit)
	var column_1 interface{}
	err := row.Scan(&column_1)
	return column_1, err
}

const getProductIDFromUser = `-- name: GetProductIDFromUser :many
SELECT scm_product.product_id from scm_account inner join scm_product on scm_account.user_id = scm_product.manufacturer_id
where scm_account.user_id = ?
`

func (q *Queries) GetProductIDFromUser(ctx context.Context, userID int32) ([]int32, error) {
	rows, err := q.db.QueryContext(ctx, getProductIDFromUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int32
	for rows.Next() {
		var product_id int32
		if err := rows.Scan(&product_id); err != nil {
			return nil, err
		}
		items = append(items, product_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProductName = `-- name: UpdateProductName :exec
Update scm_product set product_name=? where product_id=?
`

type UpdateProductNameParams struct {
	ProductName *string `json:"productName"`
	ProductID   int32   `json:"productID"`
}

func (q *Queries) UpdateProductName(ctx context.Context, arg UpdateProductNameParams) error {
	_, err := q.db.ExecContext(ctx, updateProductName, arg.ProductName, arg.ProductID)
	return err
}

const updateProductUnit = `-- name: UpdateProductUnit :exec
Update scm_product set product_unit=? where product_id=?
`

type UpdateProductUnitParams struct {
	ProductUnit *string `json:"productUnit"`
	ProductID   int32   `json:"productID"`
}

func (q *Queries) UpdateProductUnit(ctx context.Context, arg UpdateProductUnitParams) error {
	_, err := q.db.ExecContext(ctx, updateProductUnit, arg.ProductUnit, arg.ProductID)
	return err
}
