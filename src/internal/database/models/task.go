package models

import (
)

type Task struct {
	ID          int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Judul	   string `gorm:"size:255;not null" json:"judul"`
	Deskripsi  string `gorm:"size:255;not null" json:"deskripsi"`
	Status 		string `gorm:"size:50;not null" json:"status"`
	Prioritas  string `gorm:"size:50;not null" json:"prioritas"`
	AssignedTo int64  `gorm:"not null" json:"assigned_to"`
	ProjectID  int64  `gorm:"not null" json:"project_id"`
	CreatedBy  int64  `gorm:"not null" json:"created_by"`
	CreatedAt  string `gorm:"not null" json:"created_at"`
}

func (Task) TableName() string {
	return "tasks"
}