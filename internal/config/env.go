package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// LoadConfig loads the configuration from the .env file
func LoadConfig() *viper.Viper {
	config := viper.New()
	config.SetConfigFile(".env")
	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error loading config file: %w", err))
	}

	return config
}
