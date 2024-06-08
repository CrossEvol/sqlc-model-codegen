-- name: GetDetailById :one
SELECT *
FROM `Detail`
WHERE id = ?
LIMIT 1;

-- name: GetDetailByTodoId :one
SELECT *
FROM `Detail`
WHERE todo_id = ?
LIMIT 1;

-- name: GetDetails :many
SELECT *
FROM `Detail`;

-- name: GetDetailsByIds :many
SELECT *
FROM `Detail`
WHERE id IN (sqlc.slice('ids'));

-- name: GetDetailsByTodoIds :many
SELECT *
FROM `Detail`
WHERE todo_id IN (sqlc.slice('ids'));

-- name: CountDetails :one
SELECT count(*)
FROM `Detail`;

-- name: CreateDetail :execresult
INSERT INTO `Detail` (`desc`, `img_url`, `todo_id`)
VALUES (?, ?, ?);

-- name: UpdateDetail :execresult
UPDATE `Detail`
SET `desc`    = CASE WHEN @desc IS NOT NULL THEN @desc ELSE `desc` END,
    `img_url` = CASE WHEN @img_url IS NOT NULL THEN @img_url ELSE `img_url` END,
    `todo_id` = CASE WHEN @todo_id IS NOT NULL THEN @todo_id ELSE `todo_id` END
WHERE id = ?;

-- name: DeleteDetail :exec
DELETE
FROM `Detail`
WHERE id = ?;
