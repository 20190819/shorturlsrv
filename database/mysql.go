package database

import (
	"fmt"
	"log"
	"shorturlsrv/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlDB struct {
	*gorm.DB
}

var MysqlClient *mysqlDB
var MysqlClients map[string]*mysqlDB

func InitDB() error {
	MysqlClients = make(map[string]*mysqlDB)
	db, err := gorm.Open(mysql.Open(config.MysqlDefault.Dns), &config.MysqlDefault.Config)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("Conn mysql database success.")
	MysqlClient = &mysqlDB{db}
	MysqlClients[config.MysqlDefault.Name] = MysqlClient
	return nil
}

// 自定义分页
func (db *mysqlDB) Paging(page, limit int) *gorm.DB {
	offs := (page - 1) * limit
	return MysqlClient.Offset(offs).Limit(limit)
}
