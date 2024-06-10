package dto

import (
	"github.com/crossevol/sqlc-model-codegen/internal/database/sqliteDao"
	"time"
)

type PostDTO struct {
	ID          int64  `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	CreatedAt   int64  `db:"created_at" json:"created_at"`
	UpdatedAt   int64  `db:"updated_at" json:"updated_at"`
	CreatedByID string `db:"created_by_id" json:"created_by_id"`
}

type CreatePostDTO struct {
	Name        string `db:"name" json:"name"`
	CreatedAt   int64  `db:"created_at" json:"created_at"`
	UpdatedAt   int64  `db:"updated_at" json:"updated_at"`
	CreatedByID string `db:"created_by_id" json:"created_by_id"`
}

type UpdatePostDTO struct {
	Name        *string `db:"name" json:"name"`
	CreatedAt   int64   `db:"created_at" json:"created_at"`
	UpdatedAt   int64   `db:"updated_at" json:"updated_at"`
	CreatedByID *string `db:"created_by_id" json:"created_by_id"`
	ID          int64   `db:"id" json:"id"`
}

func (dto *CreatePostDTO) Map2CreatePostParams() *sqliteDao.CreatePostParams {
	CreatedAt := time.UnixMilli(dto.CreatedAt)
	UpdatedAt := time.UnixMilli(dto.UpdatedAt)

	params := sqliteDao.CreatePostParams{
		Name:        dto.Name,
		CreatedAt:   &CreatedAt,
		UpdatedAt:   &UpdatedAt,
		CreatedByID: dto.CreatedByID,
	}
	return &params
}

func (dto *UpdatePostDTO) Map2UpdatePostParams() *sqliteDao.UpdatePostParams {
	CreatedAt := time.UnixMilli(dto.CreatedAt)
	UpdatedAt := time.UnixMilli(dto.UpdatedAt)

	params := sqliteDao.UpdatePostParams{
		Name:        dto.Name,
		CreatedAt:   &CreatedAt,
		UpdatedAt:   &UpdatedAt,
		CreatedByID: dto.CreatedByID,
		ID:          dto.ID,
	}
	return &params
}

func Map2PostDTO(entity *sqliteDao.Post) *PostDTO {
	dto := PostDTO{
		ID:          entity.ID,
		Name:        entity.Name,
		CreatedAt:   entity.CreatedAt.UnixMilli(),
		UpdatedAt:   entity.UpdatedAt.UnixMilli(),
		CreatedByID: entity.CreatedByID,
	}
	return &dto
}
