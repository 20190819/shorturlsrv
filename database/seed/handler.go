package seed

import (
	"shorturlsrv/database"
	"shorturlsrv/models"
)

func Seeder() {
	secretSeeder()
	// userSeeder()
}
func secretSeeder() {
	database.Truncate("secrets")
	sec := models.SecretModel{
		Key:    "test",
		Secret: "1234567890",
	}
	database.MysqlClient.FirstOrCreate(&sec)
}

func userSeeder() {
	database.Truncate("users")
	user := models.UserModel{
		Email:    "1234567890@qq.com",
		Password: "123456",
	}
	database.MysqlClient.FirstOrCreate(&user)
}
