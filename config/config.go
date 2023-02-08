package config

import (
	"github.com/spf13/viper"
)

func ReadConfigs() {
	viper.SetConfigType("env")
	viper.SetConfigFile("./test.env")
	viper.ReadInConfig()
}

func GetConfigAsString(key string) string {
	return viper.GetString(key)
}

func GetConfigAsInt(key string) int {
	return viper.GetInt(key)
}
