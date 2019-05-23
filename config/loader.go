package config

import (
	"github.com/geekymedic/neonx/errors"
	"github.com/spf13/viper"
)

func Load() error {

	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()

	if err != nil {
		return errors.By(err)
	}

	return nil
}
