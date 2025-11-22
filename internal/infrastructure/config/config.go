package config

import (
	"fmt"
	"os"
	"time"

	"github.com/goNiki/Subscription-service/internal/domain/errorapp"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	DB     DBConfig     `yaml:"database"`
}

type ServerConfig struct {
	Env         string        `yaml:"env"`
	Host        string        `yaml:"host"`
	Port        string        `yaml:"port"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

type DBConfig struct {
	Host              string        `yaml:"host"`
	Port              string        `yaml:"port"`
	User              string        `yaml:"user"`
	Password          string        `yaml:"password"`
	Name              string        `yaml:"name"`
	Sslmode           string        `yaml:"sslmode"`
	MaxConns          int32         `yaml:"maxconns"`
	MinConns          int32         `yaml:"minconns"`
	MaxConnLifeTime   time.Duration `yaml:"maxconnlifetime"`
	MaxConnIdleTime   time.Duration `yaml:"maxconnidletime"`
	HealthCheckPeriod time.Duration `yaml:"healthcheckperiod"`
}

func InitConfig(path string) (*Config, error) {

	date, err := os.ReadFile(path)
	if err != nil {
		return &Config{}, fmt.Errorf("%w: %v", errorapp.ErrInitConfig, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(date, &cfg); err != nil {
		return &Config{}, fmt.Errorf("%w: %v", errorapp.ErrInitConfig, err)
	}

	return &cfg, nil
}
