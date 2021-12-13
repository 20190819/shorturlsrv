package database

import "shorturlsrv/models"

func Seed() {
	sec := models.Secret{
		Key:    "test",
		Secret: "1234567890",
	}

	MysqlClient.Create(&sec)
}
