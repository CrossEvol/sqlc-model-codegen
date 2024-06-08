-- name: GetTodo :one
SELECT *
FROM `Todo`
WHERE id = ?
LIMIT 1;

-- name: GetTodos :many
SELECT *
FROM `Todo`;

-- name: GetTodosByIds :many
SELECT *
FROM `Todo`
WHERE id IN (sqlc.slice('ids'));

-- name: CountTodos :one
SELECT count(*)
FROM `Todo`;

-- name: CreateTodo :execresult
INSERT INTO `Todo` (`title`, `score`, `amount`, `status`, `created_at`, `updated_at`, `deadline`, `priority`, `tags`,
                    `content`, `created_by`, `assignee_email`, `detail_id`, `category_id`)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateTodo :execresult
UPDATE `Todo`
SET `title`          = CASE WHEN @title IS NOT NULL THEN @title ELSE `title` END,
    `score`          = CASE WHEN @score IS NOT NULL THEN @score ELSE `score` END,
    `amount`         = CASE WHEN @amount IS NOT NULL THEN @amount ELSE `amount` END,
    `status`         = CASE WHEN @status IS NOT NULL THEN @status ELSE `status` END,
    `created_at`     = CASE WHEN @created_at IS NOT NULL THEN @created_at ELSE `created_at` END,
    `updated_at`     = CASE WHEN @updated_at IS NOT NULL THEN @updated_at ELSE `updated_at` END,
    `deadline`       = CASE WHEN @deadline IS NOT NULL THEN @deadline ELSE `deadline` END,
    `priority`       = CASE WHEN @priority IS NOT NULL THEN @priority ELSE `priority` END,
    `tags`           = CASE WHEN @tags IS NOT NULL THEN @tags ELSE `tags` END,
    `content`        = CASE WHEN @content IS NOT NULL THEN @content ELSE `content` END,
    `created_by`     = CASE WHEN @created_by IS NOT NULL THEN @created_by ELSE `created_by` END,
    `assignee_email` = CASE WHEN @assignee_email IS NOT NULL THEN @assignee_email ELSE `assignee_email` END,
    `detail_id`      = CASE WHEN @detail_id IS NOT NULL THEN @detail_id ELSE `detail_id` END,
    `category_id`    = CASE WHEN @category_id IS NOT NULL THEN @category_id ELSE `category_id` END
WHERE id = ?;

-- name: DeleteTodo :exec
DELETE
FROM `Todo`
WHERE id = ?;
