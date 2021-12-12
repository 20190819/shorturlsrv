package models

import (
	"time"

	"gorm.io/gorm"
)

type Urls struct {
	gorm.Model
	Url       string    `gorm:"column:url;type:varchar(1000);not null;comment:跳转链接"`
	Code      string    `gorm:"column:code;type:char(20);not null;comment:短链标识"`
	ExpiredAt time.Time `gorm:"column:expired_at;type:dateTime;comment:有效时间"`
	UserId    uint64    `gorm:"column:user_id;type:int;comment:操作用户id"`
}
