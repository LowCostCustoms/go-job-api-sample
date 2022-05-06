package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"sync"
)

var dotenvOnce sync.Once

type DatabaseConfig struct {
	DatabaseUrl        string `env:"DATABASE_URL,required"`
	MigrationsLocation string `env:"MIGRATIONS_LOCATION" envDefault:"file://migrations"`
}

type ServerConfig struct {
	Port uint16 `env:"PORT" envDefault:"8080"`
}

func Parse[T any]() (*T, error) {
	dotenvOnce.Do(func() {
		dotenvFileName := ".env"
		if err := godotenv.Load(dotenvFileName); err != nil {
			logrus.Warnf("failed to load configuration from %s: %v", dotenvFileName, err)
		}
	})

	var config T
	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func GetDatabaseConfig() (*DatabaseConfig, error) {
	return Parse[DatabaseConfig]()
}

func GetServerConfig() (*ServerConfig, error) {
	return Parse[ServerConfig]()
}
