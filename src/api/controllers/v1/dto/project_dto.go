package dto

import (
	"jk-api/internal/database/models"
	"time"
)

type CreateProjectDto struct {
	Nama             string `json:"nama"`
	Kode			 string `json:"kode"`
	Deskripsi        string `json:"deskripsi"`
	StartDate        time.Time `json:"start_date"`
	EndDate          time.Time `json:"end_date"`
}

type UpdateProjectDto struct {
	Nama             *string `json:"nama"`
	Kode             *string `json:"kode"`
	Deskripsi        *string `json:"deskripsi"`
	StartDate       *time.Time `json:"start_date"`
	EndDate         *time.Time `json:"end_date"`
}

type ProjectResponseDto struct {
	models.Project
}
