package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port          string
	Host          string
	DBPort        string
	User          string
	Password      string
	DBname        string
	BinanceAPIURL string
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	dbConfig := Config{
		Port:          ":" + viper.GetString("PORT"),
		Host:          viper.GetString("DB_HOST"),
		DBPort:        viper.GetString("DB_PORT"),
		User:          viper.GetString("DB_USER"),
		Password:      viper.GetString("DB_PASS"),
		DBname:        viper.GetString("DB_NAME"),
		BinanceAPIURL: viper.GetString("BINANCE_API_URL"),
	}

	return &dbConfig, nil
}
