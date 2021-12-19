package models

import (
	"time"

	"gorm.io/gorm"
)

type Id struct {
	ID uint `json:"id" gorm:"primarykey"`
}

type TimeAt struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
