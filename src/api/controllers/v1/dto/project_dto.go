package dto

import (
	"jk-api/internal/database/models"
)

type CreateProjectDto struct {
	Nama             string `json:"nama"`
	Kode			 string `json:"kode"`
	Deskripsi        string `json:"deskripsi"`
	StartDate        string `json:"start_date"`
	EndDate          string `json:"end_date"`
}

type UpdateProjectDto struct {
	Nama             *string `json:"nama"`
	Kode             *string `json:"kode"`
	Deskripsi        *string `json:"deskripsi"`
	StartDate       *string `json:"start_date"`
	EndDate         *string `json:"end_date"`
}

type ProjectResponseDto struct {
	models.Project
}
