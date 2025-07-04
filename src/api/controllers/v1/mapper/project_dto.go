package mapper

import (
	"jk-api/api/controllers/v1/dto"
	"jk-api/internal/database/models"
)

func CreateProjectDtoToModel(dto *dto.CreateProjectDto) (*models.Project, error) {
	if dto == nil {
		return nil, nil
	}

	data := &models.Project{
		Nama:      dto.Nama,
		Kode:      dto.Kode,
		Deskripsi: dto.Deskripsi,
		StartDate: dto.StartDate,
		EndDate:   dto.EndDate,
	}

	return data, nil
}

func UpdateProjectDtoToModel(dto *dto.UpdateProjectDto) (*models.Project, error) {
	if dto == nil {
		return nil, nil
	}

	data := &models.Project{}

	if dto.Nama != nil {
		data.Nama = *dto.Nama
	}
	if dto.Kode != nil {
		data.Kode = *dto.Kode
	}
	if dto.Deskripsi != nil {
		data.Deskripsi = *dto.Deskripsi
	}
	if dto.StartDate != nil {
		data.StartDate = *dto.StartDate
	}
	if dto.EndDate != nil {
		data.EndDate = *dto.EndDate
	}

	return data, nil
}

func ProjectModelToResponseDto(data *models.Project) (*dto.ProjectResponseDto, error) {
	if data == nil {
		return nil, nil
	}

	responseDto := &dto.ProjectResponseDto{
		Project: *data,
	}

	return responseDto, nil
}
