package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort    string
	MongoURI      string
	MongoDatabase string
	MongoTimeout  time.Duration
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("MONGO_URI", "mongodb://root:root@mongo:27017")
	viper.SetDefault("MONGO_DATABASE", "go-example")
	viper.SetDefault("MONGO_TIMEOUT", 10)

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: .env file not found or invalid: %v", err)
	}

	return &Config{
		ServerPort:    viper.GetString("SERVER_PORT"),
		MongoURI:      viper.GetString("MONGO_URI"),
		MongoDatabase: viper.GetString("MONGO_DATABASE"),
		MongoTimeout:  time.Duration(viper.GetInt("MONGO_TIMEOUT")) * time.Second,
	}
}
