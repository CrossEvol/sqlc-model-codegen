// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: Post.sql

package sqliteDao

import (
	"context"
	"database/sql"
	"strings"
	"time"
)

const CountPosts = `-- name: CountPosts :one
SELECT count(*)
FROM Post
`

func (q *Queries) CountPosts(ctx context.Context) (int64, error) {
	row := q.queryRow(ctx, q.countPostsStmt, CountPosts)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const CreatePost = `-- name: CreatePost :execresult
INSERT INTO Post (name, created_at, updated_at, created_by_id)
VALUES (?, ?, ?, ?)
`

type CreatePostParams struct {
	Name        string     `db:"name" json:"name"`
	CreatedAt   *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at"`
	CreatedByID string     `db:"created_by_id" json:"created_by_id"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (sql.Result, error) {
	return q.exec(ctx, q.createPostStmt, CreatePost,
		arg.Name,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.CreatedByID,
	)
}

const DeletePost = `-- name: DeletePost :exec
DELETE
FROM Post
WHERE id = ?
`

func (q *Queries) DeletePost(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deletePostStmt, DeletePost, id)
	return err
}

const GetPost = `-- name: GetPost :one
SELECT id, name, created_at, updated_at, created_by_id
FROM Post
WHERE id = ?
LIMIT 1
`

func (q *Queries) GetPost(ctx context.Context, id int64) (Post, error) {
	row := q.queryRow(ctx, q.getPostStmt, GetPost, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedByID,
	)
	return i, err
}

const GetPosts = `-- name: GetPosts :many
SELECT id, name, created_at, updated_at, created_by_id
FROM Post
`

func (q *Queries) GetPosts(ctx context.Context) ([]Post, error) {
	rows, err := q.query(ctx, q.getPostsStmt, GetPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CreatedByID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const GetPostsByIds = `-- name: GetPostsByIds :many
SELECT id, name, created_at, updated_at, created_by_id
FROM Post
WHERE id IN (/*SLICE:ids*/?)
`

func (q *Queries) GetPostsByIds(ctx context.Context, ids []int64) ([]Post, error) {
	query := GetPostsByIds
	var queryParams []interface{}
	if len(ids) > 0 {
		for _, v := range ids {
			queryParams = append(queryParams, v)
		}
		query = strings.Replace(query, "/*SLICE:ids*/?", strings.Repeat(",?", len(ids))[1:], 1)
	} else {
		query = strings.Replace(query, "/*SLICE:ids*/?", "NULL", 1)
	}
	rows, err := q.query(ctx, nil, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CreatedByID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const UpdatePost = `-- name: UpdatePost :execresult
UPDATE Post
SET name          = CASE WHEN ? IS NOT NULL THEN ? ELSE name END,
    created_at    = CASE WHEN ? IS NOT NULL THEN ? ELSE created_at END,
    updated_at    = CASE WHEN ? IS NOT NULL THEN ? ELSE updated_at END,
    created_by_id = CASE WHEN ? IS NOT NULL THEN ? ELSE created_by_id END
WHERE id = ?
`

type UpdatePostParams struct {
	Name        interface{} `db:"name" json:"name"`
	CreatedAt   interface{} `db:"created_at" json:"created_at"`
	UpdatedAt   interface{} `db:"updated_at" json:"updated_at"`
	CreatedByID interface{} `db:"created_by_id" json:"created_by_id"`
	ID          int64       `db:"id" json:"id"`
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (sql.Result, error) {
	return q.exec(ctx, q.updatePostStmt, UpdatePost,
		arg.Name,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.CreatedByID,
		arg.ID,
	)
}
