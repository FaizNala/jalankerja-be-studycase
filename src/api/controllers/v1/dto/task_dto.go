// task_dto.go - generated by api:generate task


package dto

import (
	"jk-api/internal/database/models"
)

// CreateTaskDto is used when creating a new Task.
type CreateTaskDto struct {
	Judul			string `json:"judul"`
	Deskripsi		string `json:"deskripsi"`
	Status			string `json:"status"`
	Prioritas		string `json:"prioritas"`
	AssignedTo		int64  `json:"assigned_to"`
	ProjectID		int64  `json:"project_id"`
	CreatedBy		int64  `json:"created_by"`
}

// UpdateTaskDto is used when updating an existing Task.
type UpdateTaskDto struct {
	Judul			*string `json:"judul"`
	Deskripsi		*string `json:"deskripsi"`
	Status			*string `json:"status"`
	Prioritas		*string `json:"prioritas"`
	AssignedTo		*int64  `json:"assigned_to"`
	ProjectID		*int64  `json:"project_id"`
	CreatedBy		*int64  `json:"created_by"`
}

// TaskResponseDto represents a detailed view of Task with related data.
type TaskResponseDto struct {
	models.Task
}
