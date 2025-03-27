// Package config provides the environment variables loading, and also exposes the config struct.
package config

import (
	"github.com/spf13/viper"
)

// Config struct exposes all configuration variables that should be needed in different parts of the application.
//
// Configuration variables are in their own nested struct.
type Config struct {
	Database struct {
		Host     string `mapstructure:"DB_HOST"`
		Password string `mapstructure:"DB_PASSWORD"`
		User     string `mapstructure:"DB_USER"`
		Name     string `mapstructure:"DB_NAME"`
		Port     int    `mapstructure:"DB_PORT"`
	} `mapstructure:",squash"`

	API struct {
		Port int `mapstructure:"API_PORT"`
	} `mapstructure:",squash"`
}

// NewConfig loads environment variables from .env file. Panics if cannot read the file or if cannot unmarshal values into Config struct.
func NewConfig() *Config {
	var config Config

	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	return &config
}
