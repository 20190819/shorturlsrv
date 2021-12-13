package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	username string `gorm:"column:username;type:varchar(100);not null;comment:用户名"`
	password string `gorm:"column:password;not null;comment:密码"`
}
