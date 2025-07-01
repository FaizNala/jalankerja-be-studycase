
package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type TaskRepository interface {
	WithTx(tx *gorm.DB) TaskRepository

	InsertTask(input *models.Task) (*models.Task, error)
	UpdateTasks(input *models.Task) (*models.Task, error)
	RemoveTask(data *models.Task) (*models.Task, error)
	RemoveTaskByID(id int64) error
	FindTask() ([]models.Task, error)
	FindTaskByID(id int64) (*models.Task, error)
}
	