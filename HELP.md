# relation "`Category`" does not exist
```shell
$ sqlc generate 
# package sqliteDao
assets\queries\Category.sql:1:1: relation "`Category`" does not exist
assets\queries\Category.sql:8:1: relation "`Category`" does not exist
assets\queries\Category.sql:12:1: relation "`Category`" does not exist
assets\queries\Category.sql:17:1: relation "`Category`" does not exist
assets\queries\Category.sql:21:1: relation "`Category`" does not exist
assets\queries\Category.sql:25:1: relation "`Category`" does not exist
assets\queries\Category.sql:33:1: relation "`Category`" does not exist
```

`Category.sql`
```sqlite
-- name: GetCategory :one
SELECT *
FROM `Category`
WHERE id = ?
LIMIT 1;

-- name: GetCategories :many
SELECT *
FROM `Category`;

-- name: GetCategoriesByIds :many
SELECT *
FROM `Category`
WHERE id IN (sqlc.slice('ids'));

-- name: CountCategories :one
SELECT count(*)
FROM `Category`;

-- name: CreateCategory :execresult
INSERT INTO `Category` (`name`, `desc`, `created_at`, `updated_at`)
VALUES (?, ?, ?, ?);

-- name: UpdateCategory :execresult
UPDATE `Category`
SET `name`       = CASE WHEN @name IS NOT NULL THEN @name ELSE `name` END,
    `desc`       = CASE WHEN @desc IS NOT NULL THEN @desc ELSE `desc` END,
    `created_at` = CASE WHEN @created_at IS NOT NULL THEN @created_at ELSE `created_at` END,
    `updated_at` = CASE WHEN @updated_at IS NOT NULL THEN @updated_at ELSE `updated_at` END
WHERE id = ?;

-- name: DeleteCategory :exec
DELETE
FROM `Category`
WHERE id = ?;

```

the table name in queries/ should be same as in the migrations/