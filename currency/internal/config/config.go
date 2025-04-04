package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	TaskStartTime  TaskStartTime  `mapstructure:"task_start_time"`
	DataBaseConfig DataBaseConfig `mapstructure:"database_config"`
}

type TaskStartTime struct {
	Hour   int `mapstructure:"hour"`
	Minute int `mapstructure:"minute"`
}

type DataBaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	DBName   string `mapstructure:"dbname"`
	Password string `mapstructure:"password"`
	User     string `mapstructure:"user"`
	SSLMode  string `mapstructure:"sslmode"`
}

func LoadConfig() (*AppConfig, error) {

	v := viper.New()
	v.AddConfigPath("currency/configs")
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
