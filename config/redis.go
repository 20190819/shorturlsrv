package config

import (
	"fmt"
	"shorturlsrv/app"

	redisStd "github.com/go-redis/redis/v8"
)

const (
	DefaultNameRedis = "default"
)

type redis struct {
	Name   string
	Addr   string
	Config redisStd.Options
}
type redisAddr string

var (
	addr         = fmt.Sprintf("%s:%s", app.EnvStr("REDIS_HOST"), app.EnvStr("REDIS_PORT"))
	RedisDefault = redis{
		Name: DefaultNameRedis,
		Addr: addr,
		Config: redisStd.Options{
			Addr: addr,
			DB:   app.EnvInt("REDIS_SELECT_DB"),
		},
	}
)
