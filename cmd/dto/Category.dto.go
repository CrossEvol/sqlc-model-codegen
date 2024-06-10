package dto

import (
	"github.com/crossevol/sqlc-model-codegen/internal/database/sqliteDao"
	"time"
)

type CategoryDTO struct {
	ID        int64  `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Desc      string `db:"desc" json:"desc"`
	CreatedAt int64  `db:"created_at" json:"created_at"`
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`
}

type CreateCategoryDTO struct {
	Name      string `db:"name" json:"name"`
	Desc      string `db:"desc" json:"desc"`
	CreatedAt int64  `db:"created_at" json:"created_at"`
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`
}

type UpdateCategoryDTO struct {
	Name      *string `db:"name" json:"name"`
	Desc      *string `db:"desc" json:"desc"`
	CreatedAt int64   `db:"created_at" json:"created_at"`
	UpdatedAt int64   `db:"updated_at" json:"updated_at"`
	ID        int64   `db:"id" json:"id"`
}

func (dto *CreateCategoryDTO) Map2CreateCategoryParams() *sqliteDao.CreateCategoryParams {
	CreatedAt := time.UnixMilli(dto.CreatedAt)
	UpdatedAt := time.UnixMilli(dto.UpdatedAt)

	params := sqliteDao.CreateCategoryParams{
		Name:      dto.Name,
		Desc:      dto.Desc,
		CreatedAt: &CreatedAt,
		UpdatedAt: &UpdatedAt,
	}
	return &params
}

func (dto *UpdateCategoryDTO) Map2UpdateCategoryParams() *sqliteDao.UpdateCategoryParams {
	CreatedAt := time.UnixMilli(dto.CreatedAt)
	UpdatedAt := time.UnixMilli(dto.UpdatedAt)

	params := sqliteDao.UpdateCategoryParams{
		Name:      dto.Name,
		Desc:      dto.Desc,
		CreatedAt: &CreatedAt,
		UpdatedAt: &UpdatedAt,
		ID:        dto.ID,
	}
	return &params
}

func Map2CategoryDTO(entity *sqliteDao.Category) *CategoryDTO {
	dto := CategoryDTO{
		ID:        entity.ID,
		Name:      entity.Name,
		Desc:      entity.Desc,
		CreatedAt: entity.CreatedAt.UnixMilli(),
		UpdatedAt: entity.UpdatedAt.UnixMilli(),
	}
	return &dto
}
