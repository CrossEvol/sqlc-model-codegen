-- name: GetUser :one
SELECT *
FROM `User`
WHERE id = ?
LIMIT 1;

-- name: GetUsers :many
SELECT *
FROM `User`;

-- name: GetUsersByIds :many
SELECT *
FROM `User`
WHERE id IN (sqlc.slice('ids'));

-- name: CountUsers :one
SELECT count(*)
FROM `User`;

-- name: CreateUser :execresult
INSERT INTO `User` (`name`, `password`, `email`, `emailVerified`, `image`, `role`)
VALUES (?, ?, ?, ?, ?, ?);

-- name: UpdateUser :execresult
UPDATE `User`
SET `name`          = CASE WHEN @name IS NOT NULL THEN @name ELSE `name` END,
    `password`      = CASE WHEN @password IS NOT NULL THEN @password ELSE `password` END,
    `email`         = CASE WHEN @email IS NOT NULL THEN @email ELSE `email` END,
    `emailVerified` = CASE WHEN @emailVerified IS NOT NULL THEN @emailVerified ELSE `emailVerified` END,
    `image`         = CASE WHEN @image IS NOT NULL THEN @image ELSE `image` END,
    `role`          = CASE WHEN @role IS NOT NULL THEN @role ELSE `role` END
WHERE id = ?;

-- name: DeleteUser :exec
DELETE
FROM `User`
WHERE id = ?;
