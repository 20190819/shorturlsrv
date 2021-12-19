package models

type Secret struct {
	Id
	Key    string `gorm:"column:key;type:varchar(50);index:idx_key;unique;comment:授权key"`
	Secret string `gorm:"column:secret;varchar(100);comment:授权密钥"`
	UserId uint64 `json:"user_id" gorm:"column:user_id;type:int;comment:操作用户id"`
	TimeAt
}

type SecretModel = Secret
