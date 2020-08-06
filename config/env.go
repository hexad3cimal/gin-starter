package config

import (
	"github.com/spf13/viper"
)

func GetEnvValue(key string) string {

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Error(err)
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Error("Invalid type assertion")
	}

	return value
}
