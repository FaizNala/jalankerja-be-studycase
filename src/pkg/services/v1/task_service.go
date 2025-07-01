package services

import (
	"errors"
	"jk-api/internal/database/config"
	"jk-api/internal/database/models"
	"jk-api/pkg/repository/adapter/sql"
	"time"

	"gorm.io/gorm"
)

type TaskService interface {
	WithTx(tx *gorm.DB) TaskService

	CreateTask(input *models.Task) (*models.Task, error)
	UpdateTask(id int64, input *models.Task) (*models.Task, error)
	DeleteTask(id int64) error
	GetAllTasks() ([]models.Task, error)
	GetTaskByID(id int64) (*models.Task, error)
	GetDB() *gorm.DB
}

type taskService struct {
	repo sql.TaskRepository
	tx   *gorm.DB
}

func NewTaskService(repo sql.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) WithTx(tx *gorm.DB) TaskService {
	return &taskService{
		repo: s.repo.WithTx(tx),
		tx:   tx,
	}
}

func (s *taskService) GetDB() *gorm.DB {
	if s.tx != nil {
		return s.tx
	}
	return config.DB
}

func (s *taskService) CreateTask(input *models.Task) (*models.Task, error) {
	input.CreatedAt = time.Now().Format(time.RFC3339)
	return s.repo.InsertTask(input)
}

func (s *taskService) UpdateTask(id int64, input *models.Task) (*models.Task, error) {
	existing, err := s.repo.FindTaskByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("task not found")
		}
		return nil, err
	}

	if input.Judul != "" {
		existing.Judul = input.Judul
	}
	if input.Deskripsi != "" {
		existing.Deskripsi = input.Deskripsi
	}
	if input.Status != "" {
		existing.Status = input.Status
	}
	if input.Prioritas != "" {
		existing.Prioritas = input.Prioritas
	}
	if input.AssignedTo != 0 {
		existing.AssignedTo = input.AssignedTo
	}
	if input.ProjectID != 0 {
		existing.ProjectID = input.ProjectID
	}
	if input.CreatedBy != 0 {
		existing.CreatedBy = input.CreatedBy
	}
	if input.CreatedAt != "" {
		existing.CreatedAt = input.CreatedAt
	}

	return s.repo.UpdateTasks(existing)
}

func (s *taskService) DeleteTask(id int64) error {
	return s.repo.RemoveTaskByID(id)
}

func (s *taskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.FindTask()
}

func (s *taskService) GetTaskByID(id int64) (*models.Task, error) {
	return s.repo.FindTaskByID(id)
}
