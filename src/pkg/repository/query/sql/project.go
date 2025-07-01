
package sql

import (
	"jk-api/internal/database/config"
	"jk-api/internal/database/models"
	"jk-api/pkg/repository/adapter/sql"

	"gorm.io/gorm"
)

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository() sql.ProjectRepository {
	return &projectRepository{db: config.DB}
}

func (repo *projectRepository) WithTx(tx *gorm.DB) sql.ProjectRepository {
	return &projectRepository{db: tx}
}

func (repo *projectRepository) InsertProject(data *models.Project) (*models.Project, error) {
	if err := repo.db.Create(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *projectRepository) UpdateProjects(data *models.Project) (*models.Project, error) {
	if err := repo.db.Save(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *projectRepository) RemoveProject(data *models.Project) (*models.Project, error) {
	if err := repo.db.Delete(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *projectRepository) RemoveProjectByID(id int64) error {
	return repo.db.Delete(&models.Project{}, id).Error
}

func (repo *projectRepository) FindProject() ([]models.Project, error) {
	var items []models.Project
	if err := repo.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (repo *projectRepository) FindProjectByID(id int64) (*models.Project, error) {
	var item models.Project
	if err := repo.db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
