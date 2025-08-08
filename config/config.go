package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string
	}
	Postgres struct {
		Host     string
		Port     int
		User     string
		Password string
		DBname   string
	}
}

func InitConfig() (*Config, error) {

	viper.SetConfigFile("/root/config/config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file err: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct err: %w", err)
	}

	return &cfg, nil

}
