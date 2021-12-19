package models

import (
	"time"
)

type url struct {
	Id
	Url       string    `json:"url" gorm:"column:url;type:varchar(1000);not null;comment:跳转链接"`
	Code      string    `json:"code" gorm:"column:code;type:char(20);not null;comment:短链标识"`
	ExpiredAt time.Time `json:"expired_at" gorm:"column:expired_at;type:dateTime;comment:有效时间"`
	UserId    uint64    `json:"user_id" gorm:"column:user_id;type:int;comment:操作用户id"`
	TimeAt
}

type UrlModel = url
