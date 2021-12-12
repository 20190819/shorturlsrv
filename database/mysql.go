package database

import (
	"fmt"
	"shorturlsrv/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlClient *gorm.DB
var MysqlClients map[string]*gorm.DB

func InitDB() error {
	MysqlClients = make(map[string]*gorm.DB)
	db, err := gorm.Open(mysql.Open(config.MysqlDefault.Dns), &config.MysqlDefault.Config)
	if err != nil {
		return err
	}
	fmt.Println("Conn mysql database success.")
	MysqlClient = db
	MysqlClients[config.MysqlDefault.Name] = MysqlClient
	return nil
}
