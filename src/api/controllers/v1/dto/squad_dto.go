package dto

import (
	"jk-api/internal/database/models"
)

type CreateSquadDto struct {
	Name string `json:"name"`
}

type UpdateSquadDto struct {
	Name *string `json:"name"`
}

type SquadResponseDto struct {
	models.Squad
}
