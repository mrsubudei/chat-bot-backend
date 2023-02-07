package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App      `yaml:"app"`
		HTTP     `yaml:"http"`
		Postgres `yaml:"postgres"`
		Logger   `yaml:"logger"`
	}

	// App -.
	App struct {
		Name    string `yaml:"name" env:"APP_NAME"`
		Version string `yaml:"version" env:"APP_VERSION"`
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
	Postgres struct {
		PoolMax  int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		URL      string `env-required:"true" yaml:"url" env:"PG_URL"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		NameDB   string `yaml:"db_name"`
	}

	// Logger -.
	Logger struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
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
