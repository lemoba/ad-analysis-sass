package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func NewConfig(p string) *viper.Viper {
	envConf := os.Getenv("APP_ENV")
	if envConf == "" {
		envConf = p
	}

	fmt.Println("load conf file:", envConf)
	return getConfig(envConf)
}

func getConfig(p string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigFile(p)
	err := conf.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return conf
}
