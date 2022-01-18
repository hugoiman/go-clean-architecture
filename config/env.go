package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func ReadEnv() (err error) {
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return err
}
