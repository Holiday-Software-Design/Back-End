// 读取配置
package configs

import (
	"log"

	"github.com/spf13/viper"
)

var Config = viper.New()

func init() {
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	Config.AddConfigPath("./configs/configuration")
	Config.WatchConfig() // 自动将配置读入Config变量

	err := Config.ReadInConfig()
	if err != nil {
		log.Fatal("Config not found: ", err)
	}
}
