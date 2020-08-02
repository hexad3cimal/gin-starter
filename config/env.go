package config

import (
	"github.com/spf13/viper"
)

func GetEnv(key string) string {

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		Logger().Error(err)
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		Logger().Error("Invalid type assertion")
	}

	return value
}
