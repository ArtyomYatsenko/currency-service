package config

import (
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"time"
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
	Host          string `env:"DB_HOST"`
	Port          string `env:"DB_PORT"`
	DBName        string `env:"DB_NAME"`
	Password      string `env:"DB_PASSWORD"`
	User          string `env:"DB_USER"`
	SSLMode       string `env:"DB_SSLMODE"`
	DirMigrations string `env:"DB_DIR_MIGRATIONS"`
}

type HttpClient struct {
	Timeout time.Duration `mapstructure:"timeout"`
}

func LoadConfig(configPath string) (*AppConfig, error) {

	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	} // Парсим конфигурацию приложения yaml файла
	var a AppConfig
	if err := v.Unmarshal(&a); err != nil { // Читаем данные из файла конфигурации .yaml
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
