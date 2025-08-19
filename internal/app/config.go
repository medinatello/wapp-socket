package app

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Config holds all configuration for the application.
type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Fakes    FakesConfig    `mapstructure:"fakes"`
	Features FeaturesConfig `mapstructure:"features"`
}

// AppConfig holds application-specific settings.
type AppConfig struct {
	Name     string `mapstructure:"name"`
	LogLevel string `mapstructure:"log_level"`
}

// FakesConfig holds settings for the fake adapters.
type FakesConfig struct {
	Seed                 int64   `mapstructure:"seed"`
	ConnectTimeoutChance float64 `mapstructure:"connect_timeout_chance"`
	ConnectFailChance    float64 `mapstructure:"connect_fail_chance"`
	AckLatencyMs         int     `mapstructure:"ack_latency_ms"`
	ReceiveIntervalMs    int     `mapstructure:"receive_interval_ms"`
}

// FeaturesConfig holds feature flags.
type FeaturesConfig struct {
	FakeMode      bool   `mapstructure:"fake_mode"`
	EventingModel string `mapstructure:"eventing_model"`
}

// LoadConfig reads configuration from file and environment variables.
func LoadConfig(path string) (*Config, error) {
	v := viper.New()

	// Set default values
	v.SetDefault("app.name", "wapp-socket")
	v.SetDefault("app.log_level", "info")
	v.SetDefault("fakes.seed", 0)
	v.SetDefault("features.fake_mode", true)
	v.SetDefault("features.eventing_model", "bus")

	// Load from config file
	v.AddConfigPath(path)
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	// Attempt to read the main config file
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
		// It's okay if the main config is not found, we can rely on defaults and env vars
	}

	// Load local config file if it exists, to override
	v.SetConfigName("config.local")
	// MergeInConfig merges the local config on top of the base config
	_ = v.MergeInConfig()

	// Load from environment variables
	v.SetEnvPrefix("WAPP")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}
