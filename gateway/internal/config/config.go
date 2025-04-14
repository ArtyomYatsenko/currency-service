package config

import (
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"time"
)

type AppConfig struct {
	Server         Server
	DataBaseConfig DataBaseConfig
}

type Server struct {
	Address      string        `mapstructure:"address"`
	Port         string        `mapstructure:"port"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
}

type DataBaseConfig struct {
	Address  string `env:"DB_HOST"`
	Password string `env:"DB_PASSWORD"`
	NumberDB int    `env:"DB_NUMBER"`
	Port     string `env:"DB_PORT"`
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

	if err := godotenv.Load(); err != nil {
		return nil, err
	} // Парсим конфигурацию базы данных из переменных окружения
	if err := env.Parse(&a.DataBaseConfig); err != nil {
		return nil, err
	}

	return &a, nil
}
