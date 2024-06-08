-- name: GetPost :one
SELECT *
FROM Post
WHERE id = ?
LIMIT 1;

-- name: GetPosts :many
SELECT *
FROM Post;

-- name: GetPostsByIds :many
SELECT *
FROM Post
WHERE id IN (sqlc.slice('ids'));

-- name: CountPosts :one
SELECT count(*)
FROM Post;

-- name: CreatePost :execresult
INSERT INTO Post (name, created_at, updated_at, created_by_id)
VALUES (?, ?, ?, ?);

-- name: UpdatePost :execresult
UPDATE Post
SET name          = CASE WHEN @name IS NOT NULL THEN @name ELSE name END,
    created_at    = CASE WHEN @created_at IS NOT NULL THEN @created_at ELSE created_at END,
    updated_at    = CASE WHEN @updated_at IS NOT NULL THEN @updated_at ELSE updated_at END,
    created_by_id = CASE WHEN @created_by_id IS NOT NULL THEN @created_by_id ELSE created_by_id END
WHERE id = ?;

-- name: DeletePost :exec
DELETE
FROM Post
WHERE id = ?;
