package database

import (
	"fmt"
	"shorturlsrv/config"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client
var RedisClients map[string]*redis.Client

func InitRedis() error {
	RedisClients = make(map[string]*redis.Client)
	RedisClient = redis.NewClient(&config.RedisDefault.Config)
	RedisClients[config.RedisDefault.Name] = RedisClient
	fmt.Println("Conn redis success.")
	return nil
}
