package mapper

import (
	"jk-api/api/controllers/v1/dto"
	"jk-api/internal/database/models"
	"time"
)

func CreateProjectDtoToModel(dto *dto.CreateProjectDto) (*models.Project, error) {
	if dto == nil {
		return nil, nil
	}

	startDate, err := time.Parse("2006-01-02", dto.StartDate)
	if err != nil {
		return nil, err
	}

	endDate, err := time.Parse("2006-01-02", dto.EndDate)
	if err != nil {
		return nil, err
	}

	data := &models.Project{
		Nama:      dto.Nama,
		Kode:      dto.Kode,
		Deskripsi: dto.Deskripsi,
		StartDate: startDate,
		EndDate:   endDate,
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
		startDate, err := time.Parse("2006-01-02", *dto.StartDate)
		if err != nil {
			return nil, err
		}
		data.StartDate = startDate
	}
	if dto.EndDate != nil {
		endDate, err := time.Parse("2006-01-02", *dto.EndDate)
		if err != nil {
			return nil, err
		}
		data.EndDate = endDate
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
