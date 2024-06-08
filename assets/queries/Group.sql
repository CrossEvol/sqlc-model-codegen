-- name: GetGroup :one
SELECT *
FROM `Group`
WHERE id = ?
LIMIT 1;

-- name: GetGroups :many
SELECT *
FROM `Group`;

-- name: GetGroupsByIds :many
SELECT *
FROM `Group`
WHERE id IN (sqlc.slice('ids'));

-- name: CountGroups :one
SELECT count(*)
FROM `Group`;

-- name: CreateGroup :execresult
INSERT INTO `Group` (`name`, `desc`, `created_at`, `updated_at`)
VALUES (?, ?, ?, ?);

-- name: UpdateGroup :execresult
UPDATE `Group`
SET `name`       = CASE WHEN @name IS NOT NULL THEN @name ELSE `name` END,
    `desc`       = CASE WHEN @desc IS NOT NULL THEN @desc ELSE `desc` END,
    `created_at` = CASE WHEN @created_at IS NOT NULL THEN @created_at ELSE `created_at` END,
    `updated_at` = CASE WHEN @updated_at IS NOT NULL THEN @updated_at ELSE `updated_at` END
WHERE id = ?;

-- name: DeleteGroup :exec
DELETE
FROM `Group`
WHERE id = ?;
