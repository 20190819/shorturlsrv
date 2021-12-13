package app

import "github.com/spf13/viper"

func IsLocal() bool {
	return Environment("local")
}

func Environment(env string) bool {
	if viper.GetString("APP_ENV") == env {
		return true
	} else {
		return false
	}
}

func EnvStr(key string) string {
	return viper.GetString(key)
}

func EnvInt(key string) int {
	return viper.GetInt(key)
}
