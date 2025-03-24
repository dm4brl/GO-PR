package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Address string
	}
	DB struct {
		Host     string
		User     string
		Password string
	}
}

func Load() *Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	var cfg Config
	viper.Unmarshal(&cfg)
	return &cfg
}
