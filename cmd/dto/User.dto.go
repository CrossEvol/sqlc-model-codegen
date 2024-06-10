package dto

import (
	"github.com/crossevol/sqlc-model-codegen/internal/database/sqliteDao"
	"time"
)

type UserDTO struct {
	ID            string  `db:"id" json:"id"`
	Name          *string `db:"name" json:"name"`
	Password      *string `db:"password" json:"password"`
	Email         *string `db:"email" json:"email"`
	EmailVerified int64   `db:"email_verified" json:"email_verified"`
	Image         *string `db:"image" json:"image"`
	Role          string  `db:"role" json:"role"`
}

type CreateUserDTO struct {
	Name          *string `db:"name" json:"name"`
	Password      *string `db:"password" json:"password"`
	Email         *string `db:"email" json:"email"`
	EmailVerified int64   `db:"email_verified" json:"email_verified"`
	Image         *string `db:"image" json:"image"`
	Role          string  `db:"role" json:"role"`
}

type UpdateUserDTO struct {
	Name          *string `db:"name" json:"name"`
	Password      *string `db:"password" json:"password"`
	Email         *string `db:"email" json:"email"`
	EmailVerified int64   `db:"email_verified" json:"email_verified"`
	Image         *string `db:"image" json:"image"`
	Role          *string `db:"role" json:"role"`
	ID            string  `db:"id" json:"id"`
}

func (dto *CreateUserDTO) Map2CreateUserParams() *sqliteDao.CreateUserParams {
	EmailVerified := time.UnixMilli(dto.EmailVerified)

	params := sqliteDao.CreateUserParams{
		Name:          dto.Name,
		Password:      dto.Password,
		Email:         dto.Email,
		EmailVerified: &EmailVerified,
		Image:         dto.Image,
		Role:          dto.Role,
	}
	return &params
}

func (dto *UpdateUserDTO) Map2UpdateUserParams() *sqliteDao.UpdateUserParams {
	EmailVerified := time.UnixMilli(dto.EmailVerified)

	params := sqliteDao.UpdateUserParams{
		Name:          dto.Name,
		Password:      dto.Password,
		Email:         dto.Email,
		EmailVerified: &EmailVerified,
		Image:         dto.Image,
		Role:          dto.Role,
		ID:            dto.ID,
	}
	return &params
}

func Map2UserDTO(entity *sqliteDao.User) *UserDTO {
	dto := UserDTO{
		ID:            entity.ID,
		Name:          entity.Name,
		Password:      entity.Password,
		Email:         entity.Email,
		EmailVerified: entity.EmailVerified.UnixMilli(),
		Image:         entity.Image,
		Role:          entity.Role,
	}
	return &dto
}
