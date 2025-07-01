package models

import (
	"time"
)

type User struct {
	ID              int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama            string         `gorm:"size:255" json:"nama"`
	Email           string         `gorm:"size:255;uniqueIndex" json:"email"`
	CreatedAt       time.Time      `json:"created_at"`
}

func (User) TableName() string {
	return "users"
}
