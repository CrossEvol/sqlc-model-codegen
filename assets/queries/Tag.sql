-- name: GetTag :one
SELECT *
FROM Tag
WHERE id = ?
LIMIT 1;

-- name: GetTags :many
SELECT *
FROM Tag;

-- name: GetTagsByParentId :many
SELECT *
FROM Tag
WHERE parent_id = ?;

-- name: GetTagsByIds :many
SELECT *
FROM Tag
WHERE id IN (sqlc.slice('ids'));

-- name: CountTags :one
SELECT count(*)
FROM Tag;

-- name: CountTagsByParentId :one
SELECT count(*)
FROM Tag
WHERE parent_id = ?;

-- name: CreateTag :execresult
INSERT INTO Tag (name, parent_id)
VALUES (?, ?);

-- name: UpdateTag :execresult
UPDATE Tag
SET name      = CASE WHEN @name IS NOT NULL THEN @name ELSE name END,
    parent_id = CASE WHEN @parent_id IS NOT NULL THEN @parent_id ELSE parent_id END
WHERE id = ?;

-- name: DeleteTag :exec
DELETE
FROM Tag
WHERE id = ?;
