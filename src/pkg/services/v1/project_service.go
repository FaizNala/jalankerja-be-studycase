package services

import (
	"errors"
	"time"
	"jk-api/internal/database/config"
	"jk-api/internal/database/models"
	"jk-api/pkg/repository/adapter/sql"
	"gorm.io/gorm"
)

type ProjectService interface {
	WithTx(tx *gorm.DB) ProjectService

	CreateProject(input *models.Project) (*models.Project, error)
	UpdateProject(id int64, input *models.Project) (*models.Project, error)
	DeleteProject(id int64) error
	GetAllProjects() ([]models.Project, error)
	GetProjectByID(id int64) (*models.Project, error)
	GetDB() *gorm.DB
}

type projectService struct {
	repo sql.ProjectRepository
	tx   *gorm.DB
}

func NewProjectService(repo sql.ProjectRepository) ProjectService {
	return &projectService{repo: repo}
}

func (s *projectService) WithTx(tx *gorm.DB) ProjectService {
	return &projectService{
		repo: s.repo.WithTx(tx),
		tx:   tx,
	}
}

func (s *projectService) GetDB() *gorm.DB {
	if s.tx != nil {
		return s.tx
	}
	return config.DB
}

func (s *projectService) CreateProject(input *models.Project) (*models.Project, error) {
	return s.repo.InsertProject(input)
}

func (s *projectService) UpdateProject(id int64, input *models.Project) (*models.Project, error) {
	existing, err := s.repo.FindProjectByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("project not found")
		}
		return nil, err
	}

	if input.Nama != "" {
		existing.Nama = input.Nama
	}
	if input.Kode != "" {
		existing.Kode = input.Kode
	}
	if input.Deskripsi != "" {
		existing.Deskripsi = input.Deskripsi
	}
	if input.StartDate != (time.Time{}) {
		existing.StartDate = input.StartDate
	}
	if input.EndDate != (time.Time{}) {
		existing.EndDate = input.EndDate
	}

	return s.repo.UpdateProjects(existing)
}

func (s *projectService) DeleteProject(id int64) error {
	return s.repo.RemoveProjectByID(id)
}

func (s *projectService) GetAllProjects() ([]models.Project, error) {
	return s.repo.FindProject()
}

func (s *projectService) GetProjectByID(id int64) (*models.Project, error) {
	return s.repo.FindProjectByID(id)
}
