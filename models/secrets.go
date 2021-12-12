package models

import "gorm.io/gorm"

type Secrets struct {
	gorm.Model
	Key    string `gorm:"column:key;type:varchar(50);index:idx_key;unique;comment:授权key"`
	Secret string `gorm:"column:secret;varchar(100);comment:授权密钥"`
}
