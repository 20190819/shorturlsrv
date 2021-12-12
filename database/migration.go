package database

import (
	"shorturlsrv/models"
)

func Migrate() {
	MysqlClient.AutoMigrate(&models.Users{})
	MysqlClient.AutoMigrate(&models.Secret{})
	MysqlClient.AutoMigrate(&models.Urls{})
}
