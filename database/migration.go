package database

import (
	"fmt"
	"shorturlsrv/models"
)

func Migrate() {
	MysqlClient.AutoMigrate(&models.UserModel{})
	MysqlClient.AutoMigrate(&models.SecretModel{})
	MysqlClient.AutoMigrate(&models.UrlModel{})
}

func Truncate(table interface{}) {
	MysqlClient.Raw(fmt.Sprintf("truncate table %s;", table))
}
