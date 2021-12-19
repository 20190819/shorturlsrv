package models

type Secret struct {
	Id
	Key    string `gorm:"column:key;type:varchar(50);index:idx_key;unique;comment:授权key"`
	Secret string `gorm:"column:secret;varchar(100);comment:授权密钥"`
	TimeAt
}

type SecretModel = Secret
