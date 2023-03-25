package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var Configuration *viper.Viper

func InitConfig() {
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("fatal error config file: %w", err))
	}

	Configuration = viper.GetViper()
}
