
package sql

import (
	"jk-api/internal/database/config"
	"jk-api/internal/database/models"
	"jk-api/pkg/repository/adapter/sql"

	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository() sql.TaskRepository {
	return &taskRepository{db: config.DB}
}

func (repo *taskRepository) WithTx(tx *gorm.DB) sql.TaskRepository {
	return &taskRepository{db: tx}
}

func (repo *taskRepository) InsertTask(data *models.Task) (*models.Task, error) {
	if err := repo.db.Create(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *taskRepository) UpdateTasks(data *models.Task) (*models.Task, error) {
	if err := repo.db.Save(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *taskRepository) RemoveTask(data *models.Task) (*models.Task, error) {
	if err := repo.db.Delete(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *taskRepository) RemoveTaskByID(id int64) error {
	return repo.db.Delete(&models.Task{}, id).Error
}

func (repo *taskRepository) FindTask() ([]models.Task, error) {
	var items []models.Task
	if err := repo.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (repo *taskRepository) FindTaskByID(id int64) (*models.Task, error) {
	var item models.Task
	if err := repo.db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
