package dto

import (
	"jk-api/internal/database/models"
)


type CreateUserDto struct {
	Nama             string `json:"nama"`
	Email            string `json:"email"`
}

type UpdateUserDto struct {
	Nama             *string `json:"nama"`
	Email            *string `json:"email"`
}

type UserResponseDto struct {
	models.User
}
