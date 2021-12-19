package models

import (
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Id
	Email    string `json:"email" gorm:"column:email;type:varchar(50);not null;uniqure;comment:邮箱"`
	Password string `json:"password" gorm:"column:password;not null;comment:密码"`
	TimeAt
}

type UserModel = user

func (us UserModel) Hash(pwd string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	if err != nil {
		return "", err
	} else {
		return string(b), nil
	}
}

func (us UserModel) CheckPwd(hasPwd, pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hasPwd), []byte(pwd))
}
