package bootstrap

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env") //  env 类型
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println(".env 配置文件未找到")
			return
		}
	}
	fmt.Println("Load config file success")
}
