package config

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

// ServerConfig holds HTTP server configuration
type ServerConfig struct {
	GRPCPort int `env:"grpcPort,required"`
}

// Config entire app configuration
type Config struct {
	Server ServerConfig
}

// Load loads service config from environment variables
//
// Returns:
//
//	error if one of required variables is not set
func Load() (*Config, error) {
	_ = godotenv.Load()

	var serverConfig ServerConfig
	if err := env.Parse(&serverConfig); err != nil {
		return nil, fmt.Errorf("config load failed: %w", err)
	}
	return &Config{
		Server: serverConfig,
	}, nil
}
