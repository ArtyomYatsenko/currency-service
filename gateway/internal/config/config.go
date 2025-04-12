package config

import (
	"github.com/spf13/viper"
	"time"
)

type AppConfig struct {
	Server Server
}
type Server struct {
	Address      string        `mapstructure:"address"`
	Port         string        `mapstructure:"port"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
}

func LoadConfig(configPath string) (*AppConfig, error) {

	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var a AppConfig

	if err := v.Unmarshal(&a); err != nil {
		return nil, err
	}

	return &a, nil
}
