package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()

	if err != nil {
		panic("Read Config Error: " + err.Error())
	}
	fmt.Println(viper.GetString("server.port"))
}
