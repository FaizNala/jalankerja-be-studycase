package models

import (
	"time"
)

type Project struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama        string    `gorm:"size:255;not null" json:"nama"`
	Kode        string    `gorm:"size:50;not null;unique" json:"kode"`
	Deskripsi   string    `gorm:"size:255;not null" json:"deskripsi"`
	StartDate   time.Time `gorm:"not null" json:"start_date"`
	EndDate     time.Time `gorm:"not null" json:"end_date"`
}

func (Project) TableName() string {
	return "projects"
}
