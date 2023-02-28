package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App         AppConfig
	HealthCheck HealthCheckConfig
	Cors        CorsConfig
}

const (
	ProductionEnv  = "production"
	StagingEnv     = "staging"
	DevelopmentEnv = "development"
)

type AppConfig struct {
	Port     int
	Fizz     string
	Buzz     string
	FizzBuzz string
}

type CorsConfig struct {
	AllowedOrigins []string
}

type HealthCheckConfig struct {
	Pattern string
	Port    string
}

func NewConfig(env, configDir string) (*Config, error) {
	viper.AddConfigPath(configDir)
	viper.SetConfigType("yml")

	var appConfig Config

	if env != "" {
		viper.SetConfigName(env)
		log.Printf("Running in env [%s] ...", env)
	} else {
		viper.SetConfigName("terminal")
		log.Printf("Running in the terminal")
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&appConfig); err != nil {
		return nil, err
	}
	return &appConfig, nil
}
