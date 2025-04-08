package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	TaskStartTime  TaskStartTime  `mapstructure:"task_start_time"`
	DataBaseConfig DataBaseConfig `mapstructure:"database_config"`
	HttpClient     HttpClient     `mapstructure:"http_client"`
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

type HttpClient struct {
	Timeout int `mapstructure:"timeout"`
}

func LoadConfig(configPath string) (*AppConfig, error) {

	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var a AppConfig

	if err := v.Unmarshal(&a); err != nil {
		return nil, err
	}

	a.DataBaseConfig.DBName = v.GetString("DB_NAME")
	a.DataBaseConfig.User = v.GetString("DB_USER")
	a.DataBaseConfig.Password = v.GetString("DB_PASSWORD")

	return &a, nil

}
