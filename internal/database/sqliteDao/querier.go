// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package sqliteDao

import (
	"context"
	"database/sql"
)

type Querier interface {
	CountCategories(ctx context.Context) (int64, error)
	CountDetails(ctx context.Context) (int64, error)
	CountGroups(ctx context.Context) (int64, error)
	CountPosts(ctx context.Context) (int64, error)
	CountTags(ctx context.Context) (int64, error)
	CountTagsByParentId(ctx context.Context, parentID *int64) (int64, error)
	CountTodos(ctx context.Context) (int64, error)
	CountTodosOnGroups(ctx context.Context) (int64, error)
	CountUsers(ctx context.Context) (int64, error)
	CreateCategory(ctx context.Context, arg CreateCategoryParams) (sql.Result, error)
	CreateDetail(ctx context.Context, arg CreateDetailParams) (sql.Result, error)
	CreateGroup(ctx context.Context, arg CreateGroupParams) (sql.Result, error)
	CreatePost(ctx context.Context, arg CreatePostParams) (sql.Result, error)
	CreateTag(ctx context.Context, arg CreateTagParams) (sql.Result, error)
	CreateTodo(ctx context.Context, arg CreateTodoParams) (sql.Result, error)
	CreateTodosOnGroup(ctx context.Context, arg CreateTodosOnGroupParams) (sql.Result, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error)
	DeleteCategory(ctx context.Context, id int64) error
	DeleteDetail(ctx context.Context, id int64) error
	DeleteGroup(ctx context.Context, id int64) error
	DeletePost(ctx context.Context, id int64) error
	DeleteTag(ctx context.Context, id int64) error
	DeleteTodo(ctx context.Context, id int64) error
	DeleteTodosOnGroup(ctx context.Context, arg DeleteTodosOnGroupParams) error
	DeleteUser(ctx context.Context, id string) error
	GetCategories(ctx context.Context) ([]Category, error)
	GetCategoriesByIds(ctx context.Context, ids []int64) ([]Category, error)
	GetCategory(ctx context.Context, id int64) (Category, error)
	GetDetailById(ctx context.Context, id int64) (Detail, error)
	GetDetailByTodoId(ctx context.Context, todoID int64) (Detail, error)
	GetDetails(ctx context.Context) ([]Detail, error)
	GetDetailsByIds(ctx context.Context, ids []int64) ([]Detail, error)
	GetDetailsByTodoIds(ctx context.Context, ids []int64) ([]Detail, error)
	GetGroup(ctx context.Context, id int64) (Group, error)
	GetGroups(ctx context.Context) ([]Group, error)
	GetGroupsByIds(ctx context.Context, ids []int64) ([]Group, error)
	GetPost(ctx context.Context, id int64) (Post, error)
	GetPosts(ctx context.Context) ([]Post, error)
	GetPostsByIds(ctx context.Context, ids []int64) ([]Post, error)
	GetTag(ctx context.Context, id int64) (Tag, error)
	GetTags(ctx context.Context) ([]Tag, error)
	GetTagsByIds(ctx context.Context, ids []int64) ([]Tag, error)
	GetTagsByParentId(ctx context.Context, parentID *int64) ([]Tag, error)
	GetTodo(ctx context.Context, id int64) (Todo, error)
	GetTodos(ctx context.Context) ([]Todo, error)
	GetTodosByIds(ctx context.Context, ids []int64) ([]Todo, error)
	GetTodosOnGroupByGroupId(ctx context.Context, groupID int64) (TodosOnGroup, error)
	GetTodosOnGroupByTodoId(ctx context.Context, todoID int64) (TodosOnGroup, error)
	GetTodosOnGroups(ctx context.Context) ([]TodosOnGroup, error)
	GetTodosOnGroupsByGroupIds(ctx context.Context, groupIds []int64) ([]TodosOnGroup, error)
	GetTodosOnGroupsByTodoIds(ctx context.Context, todoIds []int64) ([]TodosOnGroup, error)
	GetUser(ctx context.Context, id string) (User, error)
	GetUsers(ctx context.Context) ([]User, error)
	GetUsersByIds(ctx context.Context, ids []string) ([]User, error)
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (sql.Result, error)
	UpdateDetail(ctx context.Context, arg UpdateDetailParams) (sql.Result, error)
	UpdateGroup(ctx context.Context, arg UpdateGroupParams) (sql.Result, error)
	UpdatePost(ctx context.Context, arg UpdatePostParams) (sql.Result, error)
	UpdateTag(ctx context.Context, arg UpdateTagParams) (sql.Result, error)
	UpdateTodo(ctx context.Context, arg UpdateTodoParams) (sql.Result, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (sql.Result, error)
}

var _ Querier = (*Queries)(nil)
