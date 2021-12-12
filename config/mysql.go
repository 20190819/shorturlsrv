package config

import (
	"shorturlsrv/app"

	"gorm.io/gorm"
)

const (
	DefaultNameMysql = "default"
)

type mysql struct {
	Name   string
	Dns    string
	Config gorm.Config
}

var (
	MysqlDefault = mysql{
		Name:   DefaultNameMysql,
		Dns:    app.EnvStr("MYSQL_DNS"),
		Config: gorm.Config{},
	}
)
