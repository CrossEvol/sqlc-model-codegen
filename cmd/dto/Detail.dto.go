package dto

import (
	"github.com/crossevol/sqlc-model-codegen/internal/database/sqliteDao"
)

type DetailDTO struct {
	ID     int64  `db:"id" json:"id"`
	Desc   string `db:"desc" json:"desc"`
	ImgUrl string `db:"img_url" json:"img_url"`
	TodoID int64  `db:"todo_id" json:"todo_id"`
}

type CreateDetailDTO struct {
	Desc   string `db:"desc" json:"desc"`
	ImgUrl string `db:"img_url" json:"img_url"`
	TodoID int64  `db:"todo_id" json:"todo_id"`
}

type UpdateDetailDTO struct {
	Desc   *string `db:"desc" json:"desc"`
	ImgUrl *string `db:"img_url" json:"img_url"`
	TodoID *int64  `db:"todo_id" json:"todo_id"`
	ID     int64   `db:"id" json:"id"`
}

func (dto *CreateDetailDTO) Map2CreateDetailParams() *sqliteDao.CreateDetailParams {

	params := sqliteDao.CreateDetailParams{
		Desc:   dto.Desc,
		ImgUrl: dto.ImgUrl,
		TodoID: dto.TodoID,
	}
	return &params
}

func (dto *UpdateDetailDTO) Map2UpdateDetailParams() *sqliteDao.UpdateDetailParams {

	params := sqliteDao.UpdateDetailParams{
		Desc:   dto.Desc,
		ImgUrl: dto.ImgUrl,
		TodoID: dto.TodoID,
		ID:     dto.ID,
	}
	return &params
}

func Map2DetailDTO(entity *sqliteDao.Detail) *DetailDTO {
	dto := DetailDTO{
		ID:     entity.ID,
		Desc:   entity.Desc,
		ImgUrl: entity.ImgUrl,
		TodoID: entity.TodoID,
	}
	return &dto
}
