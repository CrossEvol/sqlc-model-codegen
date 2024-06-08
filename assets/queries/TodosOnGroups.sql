-- name: GetTodosOnGroupByTodoId :one
SELECT *
FROM TodosOnGroups
WHERE todo_id = ?
LIMIT 1;

-- name: GetTodosOnGroupByGroupId :one
SELECT *
FROM TodosOnGroups
WHERE group_id = ?
LIMIT 1;

-- name: GetTodosOnGroups :many
SELECT *
FROM TodosOnGroups;

-- name: GetTodosOnGroupsByTodoIds :many
SELECT *
FROM TodosOnGroups
WHERE todo_id IN (sqlc.slice('todo_ids'));

-- name: GetTodosOnGroupsByGroupIds :many
SELECT *
FROM TodosOnGroups
WHERE group_id IN (sqlc.slice('group_ids'));

-- name: CountTodosOnGroups :one
SELECT count(*)
FROM TodosOnGroups;

-- name: CreateTodosOnGroup :execresult
INSERT INTO TodosOnGroups (todo_id, group_id)
VALUES (?, ?);

-- name: DeleteTodosOnGroup :exec
DELETE
FROM TodosOnGroups
WHERE todo_id = ?
  and group_id = ?;
