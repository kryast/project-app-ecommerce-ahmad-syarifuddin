package util

import "github.com/spf13/viper"

type Configuration struct {
	AppName string
	Port    string
	Debug   bool
	DB      DatabaseConfig
}

type DatabaseConfig struct {
	Name     string
	Username string
	Password string
	Host     string
}

func ReadConfiguration() (Configuration, error) {
	viper.AutomaticEnv()

	return Configuration{
		AppName: viper.GetString("APP_NAME"),
		Port:    viper.GetString("PORT"),
		Debug:   viper.GetBool("DEBUG"),
		DB: DatabaseConfig{
			Name:     viper.GetString("DATABES_NAME"),
			Username: viper.GetString("DATABASE_USERNAME"),
			Password: viper.GetString("DATABASE_PASSWORD"),
			Host:     viper.GetString("DATABASE_HOST"),
		},
	}, nil
}
