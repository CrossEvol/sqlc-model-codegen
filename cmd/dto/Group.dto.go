package dto

import (
	"github.com/crossevol/sqlc-model-codegen/internal/database/sqliteDao"
	"time"
)

type GroupDTO struct {
	ID        int64  `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Desc      string `db:"desc" json:"desc"`
	CreatedAt int64  `db:"created_at" json:"created_at"`
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`
}

type CreateGroupDTO struct {
	Name      string `db:"name" json:"name"`
	Desc      string `db:"desc" json:"desc"`
	CreatedAt int64  `db:"created_at" json:"created_at"`
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`
}

type UpdateGroupDTO struct {
	Name      *string `db:"name" json:"name"`
	Desc      *string `db:"desc" json:"desc"`
	CreatedAt int64   `db:"created_at" json:"created_at"`
	UpdatedAt int64   `db:"updated_at" json:"updated_at"`
	ID        int64   `db:"id" json:"id"`
}

func (dto *CreateGroupDTO) Map2CreateGroupParams() *sqliteDao.CreateGroupParams {
	CreatedAt := time.UnixMilli(dto.CreatedAt)
	UpdatedAt := time.UnixMilli(dto.UpdatedAt)

	params := sqliteDao.CreateGroupParams{
		Name:      dto.Name,
		Desc:      dto.Desc,
		CreatedAt: &CreatedAt,
		UpdatedAt: &UpdatedAt,
	}
	return &params
}

func (dto *UpdateGroupDTO) Map2UpdateGroupParams() *sqliteDao.UpdateGroupParams {
	CreatedAt := time.UnixMilli(dto.CreatedAt)
	UpdatedAt := time.UnixMilli(dto.UpdatedAt)

	params := sqliteDao.UpdateGroupParams{
		Name:      dto.Name,
		Desc:      dto.Desc,
		CreatedAt: &CreatedAt,
		UpdatedAt: &UpdatedAt,
		ID:        dto.ID,
	}
	return &params
}

func Map2GroupDTO(entity *sqliteDao.Group) *GroupDTO {
	dto := GroupDTO{
		ID:        entity.ID,
		Name:      entity.Name,
		Desc:      entity.Desc,
		CreatedAt: entity.CreatedAt.UnixMilli(),
		UpdatedAt: entity.UpdatedAt.UnixMilli(),
	}
	return &dto
}
