package dto

import (
	"github.com/crossevol/sqlc-model-codegen/internal/database/sqliteDao"
)

type TagDTO struct {
	ID       int64  `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	ParentID *int64 `db:"parent_id" json:"parent_id"`
}

type CreateTagDTO struct {
	Name     string `db:"name" json:"name"`
	ParentID *int64 `db:"parent_id" json:"parent_id"`
}

type UpdateTagDTO struct {
	Name     *string `db:"name" json:"name"`
	ParentID *int64  `db:"parent_id" json:"parent_id"`
	ID       int64   `db:"id" json:"id"`
}

func (dto *CreateTagDTO) Map2CreateTagParams() *sqliteDao.CreateTagParams {

	params := sqliteDao.CreateTagParams{
		Name:     dto.Name,
		ParentID: dto.ParentID,
	}
	return &params
}

func (dto *UpdateTagDTO) Map2UpdateTagParams() *sqliteDao.UpdateTagParams {

	params := sqliteDao.UpdateTagParams{
		Name:     dto.Name,
		ParentID: dto.ParentID,
		ID:       dto.ID,
	}
	return &params
}

func Map2TagDTO(entity *sqliteDao.Tag) *TagDTO {
	dto := TagDTO{
		ID:       entity.ID,
		Name:     entity.Name,
		ParentID: entity.ParentID,
	}
	return &dto
}
