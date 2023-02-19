package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
		Mailer       `yaml:"mailer"`
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
		SessionExpiringTime      int    `yaml:"session_expiring_time"`
		TokenName                string `yaml:"token_name"`
		VerificationExpiringTime int    `yaml:"verification_expiring_time"`
	}

	Mailer struct {
		SmtpHost     string `yaml:"host"`
		Port         int    `yaml:"port"`
		CallBackHost string `yaml:"call_back_host"`
		AuthEmail    string `env-required:"true" yaml:"auth_email" env:"AUTH_EMAIL"`
		AuthName     string `env-required:"true" yaml:"auth_name" env:"AUTH_NAME"`
		AuthPassword string `env-required:"true" yaml:"auth_password" env:"AUTH_PASSWORD"`
	}
)

// NewConfig returns app config.
func NewConfig(path, envPath string) (*Config, error) {
	cfg := &Config{}

	err := setEnv(envPath)
	if err != nil {
		return nil, fmt.Errorf("config - setEnv: %w", err)
	}

	err = cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}

func setEnv(path string) error {
	var file *os.File
	var err error
	file, err = os.Open(path)
	if err != nil {
		return fmt.Errorf("setEnv - Open: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				os.Setenv(key, value)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("setEnv - Scan: %w", err)
	}
	return nil
}
