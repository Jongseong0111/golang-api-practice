-- name: CreateProduct :execresult
INSERT INTO scm_product (
    product_unit, manufacturer_id, product_name, reg_user_id
) VALUES ( ?, ?, ?, ?);

-- name: UpdateProductUnit :exec
Update scm_product set product_unit=? where product_id=?;

-- name: UpdateProductName :exec
Update scm_product set product_name=? where product_id=?;

-- name: CheckDuplicateProductUnit :many
SELECT product_unit from scm_product where product_unit=?;

-- name: GetProductFromProductUnit :one
SELECT (product_id, product_unit, product_name, manufacturer_id, reg_datetime, reg_user_id, mod_datetime, mod_user_id)
from scm_product where product_unit=?;

-- name: GetProductIDFromUser :many
SELECT scm_product.product_id from scm_account inner join scm_product on scm_account.user_id = scm_product.manufacturer_id
where scm_account.user_id = ?;

-- name: GetProductFromProductID :one
SELECT * FROM scm_product
WHERE product_id = ? LIMIT 1;

-- name: GetUserFromProduct :one
SELECT scm_account.user_id from scm_product inner join scm_account
    on scm_product.manufacturer_id=scm_account.user_id where scm_product.product_id=?;

