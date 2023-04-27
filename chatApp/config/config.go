package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DB struct {
		Host     string
		Port     int
		Username string
		Password string
		Name     string
	}
	Server struct {
		Host string
		Port int
	}
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read config file: %s", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Failed to unmarshal config file: %s", err)
	}

	fmt.Printf("Loaded config: %+v\n", config)

	return config, nil
}
