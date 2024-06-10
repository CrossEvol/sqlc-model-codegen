package dto

import (
	"github.com/crossevol/sqlc-model-codegen/internal/database/sqliteDao"
	"time"
)

type TodoDTO struct {
	ID            int64   `db:"id" json:"id"`
	Title         string  `db:"title" json:"title"`
	Score         int64   `db:"score" json:"score"`
	Amount        float64 `db:"amount" json:"amount"`
	Status        string  `db:"status" json:"status"`
	CreatedAt     int64   `db:"created_at" json:"created_at"`
	UpdatedAt     int64   `db:"updated_at" json:"updated_at"`
	Deadline      int64   `db:"deadline" json:"deadline"`
	Priority      string  `db:"priority" json:"priority"`
	Tags          string  `db:"tags" json:"tags"`
	Content       string  `db:"content" json:"content"`
	CreatedBy     string  `db:"created_by" json:"created_by"`
	AssigneeEmail string  `db:"assignee_email" json:"assignee_email"`
	DetailID      *int64  `db:"detail_id" json:"detail_id"`
	CategoryID    *int64  `db:"category_id" json:"category_id"`
}

type CreateTodoDTO struct {
	Title         string  `db:"title" json:"title"`
	Score         int64   `db:"score" json:"score"`
	Amount        float64 `db:"amount" json:"amount"`
	Status        string  `db:"status" json:"status"`
	CreatedAt     int64   `db:"created_at" json:"created_at"`
	UpdatedAt     int64   `db:"updated_at" json:"updated_at"`
	Deadline      int64   `db:"deadline" json:"deadline"`
	Priority      string  `db:"priority" json:"priority"`
	Tags          string  `db:"tags" json:"tags"`
	Content       string  `db:"content" json:"content"`
	CreatedBy     string  `db:"created_by" json:"created_by"`
	AssigneeEmail string  `db:"assignee_email" json:"assignee_email"`
	DetailID      *int64  `db:"detail_id" json:"detail_id"`
	CategoryID    *int64  `db:"category_id" json:"category_id"`
}

type UpdateTodoDTO struct {
	Title         *string  `db:"title" json:"title"`
	Score         *int64   `db:"score" json:"score"`
	Amount        *float64 `db:"amount" json:"amount"`
	Status        *string  `db:"status" json:"status"`
	CreatedAt     int64    `db:"created_at" json:"created_at"`
	UpdatedAt     int64    `db:"updated_at" json:"updated_at"`
	Deadline      int64    `db:"deadline" json:"deadline"`
	Priority      *string  `db:"priority" json:"priority"`
	Tags          *string  `db:"tags" json:"tags"`
	Content       *string  `db:"content" json:"content"`
	CreatedBy     *string  `db:"created_by" json:"created_by"`
	AssigneeEmail *string  `db:"assignee_email" json:"assignee_email"`
	DetailID      *int64   `db:"detail_id" json:"detail_id"`
	CategoryID    *int64   `db:"category_id" json:"category_id"`
	ID            int64    `db:"id" json:"id"`
}

func (dto *CreateTodoDTO) Map2CreateTodoParams() *sqliteDao.CreateTodoParams {
	CreatedAt := time.UnixMilli(dto.CreatedAt)
	UpdatedAt := time.UnixMilli(dto.UpdatedAt)
	Deadline := time.UnixMilli(dto.Deadline)

	params := sqliteDao.CreateTodoParams{
		Title:         dto.Title,
		Score:         dto.Score,
		Amount:        dto.Amount,
		Status:        dto.Status,
		CreatedAt:     &CreatedAt,
		UpdatedAt:     &UpdatedAt,
		Deadline:      &Deadline,
		Priority:      dto.Priority,
		Tags:          dto.Tags,
		Content:       dto.Content,
		CreatedBy:     dto.CreatedBy,
		AssigneeEmail: dto.AssigneeEmail,
		DetailID:      dto.DetailID,
		CategoryID:    dto.CategoryID,
	}
	return &params
}

func (dto *UpdateTodoDTO) Map2UpdateTodoParams() *sqliteDao.UpdateTodoParams {
	CreatedAt := time.UnixMilli(dto.CreatedAt)
	UpdatedAt := time.UnixMilli(dto.UpdatedAt)
	Deadline := time.UnixMilli(dto.Deadline)

	params := sqliteDao.UpdateTodoParams{
		Title:         dto.Title,
		Score:         dto.Score,
		Amount:        dto.Amount,
		Status:        dto.Status,
		CreatedAt:     &CreatedAt,
		UpdatedAt:     &UpdatedAt,
		Deadline:      &Deadline,
		Priority:      dto.Priority,
		Tags:          dto.Tags,
		Content:       dto.Content,
		CreatedBy:     dto.CreatedBy,
		AssigneeEmail: dto.AssigneeEmail,
		DetailID:      dto.DetailID,
		CategoryID:    dto.CategoryID,
		ID:            dto.ID,
	}
	return &params
}

func Map2TodoDTO(entity *sqliteDao.Todo) *TodoDTO {
	dto := TodoDTO{
		ID:            entity.ID,
		Title:         entity.Title,
		Score:         entity.Score,
		Amount:        entity.Amount,
		Status:        entity.Status,
		CreatedAt:     entity.CreatedAt.UnixMilli(),
		UpdatedAt:     entity.UpdatedAt.UnixMilli(),
		Deadline:      entity.Deadline.UnixMilli(),
		Priority:      entity.Priority,
		Tags:          entity.Tags,
		Content:       entity.Content,
		CreatedBy:     entity.CreatedBy,
		AssigneeEmail: entity.AssigneeEmail,
		DetailID:      entity.DetailID,
		CategoryID:    entity.CategoryID,
	}
	return &dto
}
