package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App          `yaml:"app"`
		Postgres     `yaml:"postgres"`
		Logger       `yaml:"logger"`
		Grpc         `yaml:"grpc"`
		TokenManager `yaml:"token_manager"`
	}

	// App -.
	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}

	// Grpc -.
	Grpc struct {
		Port              string `yaml:"port"`
		MaxConnectionIdle int64  `yaml:"max_connection_idle"`
		Timeout           int64  `yaml:"timeout"`
		MaxConnectionAge  int64  `yaml:"max_connection_age"`
		Host              string `yaml:"host"`
	}

	// Postgres -.
	Postgres struct {
		PoolMax int    `yaml:"pool_max"`
		URL     string `yaml:"url"`
	}

	// Logger -.
	Logger struct {
		Level string `yaml:"log_level"`
	}

	// TokenManager -.
	TokenManager struct {
		SessionExpiringTime int    `json:"session_expiring_time"`
		TokenName           string `json:"token_name"`
	}
)

// NewConfig returns app config.
func NewConfig(path string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
