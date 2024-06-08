-- name: GetCategory :one
SELECT *
FROM Category
WHERE id = ?
LIMIT 1;

-- name: GetCategories :many
SELECT *
FROM Category;

-- name: GetCategoriesByIds :many
SELECT *
FROM Category
WHERE id IN (sqlc.slice('ids'));

-- name: CountCategories :one
SELECT count(*)
FROM Category;

-- name: CreateCategory :execresult
INSERT INTO Category (name, desc, created_at, updated_at)
VALUES (?, ?, ?, ?);

-- name: UpdateCategory :execresult
UPDATE Category
SET name       = CASE WHEN @name IS NOT NULL THEN @name ELSE name END,
    desc       = CASE WHEN @desc IS NOT NULL THEN @desc ELSE desc END,
    created_at = CASE WHEN @created_at IS NOT NULL THEN @created_at ELSE created_at END,
    updated_at = CASE WHEN @updated_at IS NOT NULL THEN @updated_at ELSE updated_at END
WHERE id = ?;

-- name: DeleteCategory :exec
DELETE
FROM Category
WHERE id = ?;
