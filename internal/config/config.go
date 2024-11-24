package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	LogLevel string
	// Add other configuration fields here
}

func Load(cfgFile string) (*Config, error) {
	cfg := &Config{
		LogLevel: viper.GetString("log_level"),
	}
	return cfg, nil
}
