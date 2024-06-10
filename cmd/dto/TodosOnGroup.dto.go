package dto

import (
	"github.com/crossevol/sqlc-model-codegen/internal/database/sqliteDao"
)

type TodosOnGroupDTO struct {
	TodoID     int64 `db:"todo_id" json:"todo_id"`
	GroupID    int64 `db:"group_id" json:"group_id"`
	AssignedAt int64 `db:"assigned_at" json:"assigned_at"`
}

type CreateTodosOnGroupDTO struct {
	TodoID  int64 `db:"todo_id" json:"todo_id"`
	GroupID int64 `db:"group_id" json:"group_id"`
}

func (dto *CreateTodosOnGroupDTO) Map2CreateTodosOnGroupParams() *sqliteDao.CreateTodosOnGroupParams {

	params := sqliteDao.CreateTodosOnGroupParams{
		TodoID:  dto.TodoID,
		GroupID: dto.GroupID,
	}
	return &params
}

func Map2TodosOnGroupDTO(entity *sqliteDao.TodosOnGroup) *TodosOnGroupDTO {
	dto := TodosOnGroupDTO{
		TodoID:     entity.TodoID,
		GroupID:    entity.GroupID,
		AssignedAt: entity.AssignedAt.UnixMilli(),
	}
	return &dto
}
