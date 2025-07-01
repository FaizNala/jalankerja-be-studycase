package handlers

import (
	"jk-api/api/controllers/v1/dto"
	"jk-api/api/controllers/v1/mapper"
	"jk-api/internal/database/models"
	"jk-api/pkg/services/v1"
)

type UserHandler struct {
	Service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) CreateUserHandler(input *dto.CreateUserDto) (*dto.UserResponseDto, error) {
	db := h.Service.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()

	userService := h.Service.WithTx(db)

	newData, err := mapper.CreateUserDtoToModel(input)
	if err != nil {
		return nil, err
	}

	createdData, err := userService.CreateUser(newData)
	if err != nil {
		db.Rollback()
		return nil, err
	}

	if err := db.Commit().Error; err != nil {
		return nil, err
	}

	return mapper.UserModelToResponseDto(createdData)
}

func (h *UserHandler) UpdateUserHandler(id int64, input *dto.UpdateUserDto) (*dto.UserResponseDto, error) {
	updateData, err := mapper.UpdateUserDtoToModel(input)
	if err != nil {
		return nil, err
	}

	updatedData, err := h.Service.UpdateUser(id, updateData)
	if err != nil {
		return nil, err
	}

	return mapper.UserModelToResponseDto(updatedData)
}

func (h *UserHandler) DeleteUserHandler(id int64) error {
	return h.Service.DeleteUser(id)
}

func (h *UserHandler) GetUserByIDHandler(id int64) (*models.User, error) {
	return h.Service.GetUserByID(id)
}

func (h *UserHandler) GetAllUsersHandler() ([]models.User, error) {
	return h.Service.GetAllUsers()
}
