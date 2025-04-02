package config

import "github.com/spf13/viper"

func LoadConfig() error {
	viper.AddConfigPath("currency/internal/config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
