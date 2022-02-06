package seed

import (
	"shorturlsrv/database"
	"shorturlsrv/models"
)

func Seeder() {
	// secretSeeder()
	userSeeder()
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
	pwd, _ := models.UserModel{}.Hash("123456")
	user := models.UserModel{
		Email:    "1234567890@qq.com",
		Password: pwd,
	}
	database.MysqlClient.FirstOrCreate(&user)
}
