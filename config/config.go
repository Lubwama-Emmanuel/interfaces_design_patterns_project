package config

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/Lubwama-Emmanuel/Interfaces/storage/mongodb"
	"github.com/Lubwama-Emmanuel/Interfaces/storage/postgres"
)

type Config struct {
	Mongo    mongodb.MongoConfig
	Postgres postgres.Config
}

func NewConfig() (Config, error) {
	// Setting the viper configuration file, name , and path
	viper.SetConfigFile(".env")

	viperErr := viper.ReadInConfig()
	if viperErr != nil {
		return Config{}, fmt.Errorf("failed to load env variables: %w", viperErr)
	}

	return Config{
		Mongo: mongodb.MongoConfig{
			URL: viper.GetString("MONGODB_URL"),
		},
		Postgres: postgres.Config{
			Host:     viper.GetString("PG_HOST"),
			Port:     viper.GetString("PG_PORT"),
			Password: viper.GetString("PG_PASSWORD"),
			User:     viper.GetString("PG_USER"),
		},
	}, nil
}
