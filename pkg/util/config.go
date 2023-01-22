package util

import (
	"github.com/spf13/viper"
	"log"
)

func SetupConfig() error {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("could not read from env file")
	}

	return nil
}
