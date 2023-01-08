package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App    `yaml:"app"`
		HTTP   `yaml:"http"`
		MySql  `yaml:"mysql"`
		Logger `yaml:"logger"`
	}

	// App -.
	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}

	// HTTP -.
	HTTP struct {
		Host            string `yaml:"host"`
		Port            string `yaml:"port"`
		ReadTimeout     int    `yaml:"read_timeout"`
		WriteTimeout    int    `yaml:"write_timeout"`
		ShutDownTimeout int    `yaml:"shutdown_timeout"`
	}

	// MySql -.
	MySql struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		TimeZone string `yaml:"time_zone"`
		Location string `yaml:"location"`
	}

	// Logger -.
	Logger struct {
		Level string `yaml:"log_level"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
