package database

import (
	"shorturlsrv/models"
)

func Migrate() {
	MysqlClient.AutoMigrate(&models.Users{})
	MysqlClient.AutoMigrate(&models.Secrets{})
	MysqlClient.AutoMigrate(&models.Urls{})
}
