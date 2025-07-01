package mapper

import (
	"jk-api/api/controllers/v1/dto"
	"jk-api/internal/database/models"
)

func CreateUserDtoToModel(dto *dto.CreateUserDto) (*models.User, error) {
	if dto == nil {
		return nil, nil
	}

	data := &models.User{
		Nama: dto.Nama,
		Email: dto.Email,
	}

	return data, nil
}

func UpdateUserDtoToModel(dto *dto.UpdateUserDto) (*models.User, error) {
	if dto == nil {
		return nil, nil
	}

	data := &models.User{}

	if dto.Nama != nil {
		data.Nama = *dto.Nama
	}
	if dto.Email != nil {
		data.Email = *dto.Email
	}

	return data, nil
}

func UserModelToResponseDto(data *models.User) (*dto.UserResponseDto, error) {
	if data == nil {
		return nil, nil
	}

	responseDto := &dto.UserResponseDto{
		User: *data,
	}

	return responseDto, nil
}
