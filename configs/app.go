package configs

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

type AppConfig struct {
	APP_NAME string
	APP_ENV  string
	APP_PORT string
	APP_URL  string

	DB_CONNECTION string
	DB_HOST       string
	DB_PORT       string
	DB_DATABASE   string
	DB_USERNAME   string
	DB_PASSWORD   string

	DB_MAX_IDLE_CONNECTION      uint8
	DB_MAX_OPEN_CONNECTION      uint8
	DB_MAX_LIFETIME_CONNECTION  uint8
	DB_MAX_IDLE_TIME_CONNECTION uint8

	JWT_SECRET string
}

var ENV *AppConfig

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	errReadConfig := viper.ReadInConfig()
	if errReadConfig != nil {
		log.Fatal("Error reading config file, ", errReadConfig)
	}

	errUnmarshal := viper.Unmarshal(&ENV)
	if errUnmarshal != nil {
		log.Fatal("Error unmarshal config file, ", errUnmarshal)
	}

	log.Info("Config loaded successfully")
}
