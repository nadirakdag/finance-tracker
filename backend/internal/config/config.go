package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server  ServerConfig  `yaml:"server"`
	Cors    CorsConfig    `yaml:"cors"`
	Logging LoggingConfig `yaml:"logging"`
	Metrics MetricsConfig `yaml:"metrics"`
}

type ServerConfig struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

type CorsConfig struct {
	AllowedOrigins []string `yaml:"allowed_origins"`
	AllowedMethods []string `yaml:"allowed_methods"`
	AllowedHeaders []string `yaml:"allowed_headers"`
}

type LoggingConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

type MetricsConfig struct {
	Enabled bool   `yaml:"enabled"`
	Path    string `yaml:"path"`
}

// Load reads the configuration from the specified file
func Load(configPath string) (*Config, error) {
	// If no config path is provided, try to find it in common locations
	if configPath == "" {
		configPath = findConfigFile()
	}

	// Read config file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	// Parse YAML
	config := &Config{}
	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, err
	}

	// Set defaults if needed
	setDefaults(config)

	return config, nil
}

// findConfigFile looks for the config file in common locations
func findConfigFile() string {
	locations := []string{
		"configs/config.yaml",
		"config.yaml",
		"../configs/config.yaml",
		"../../configs/config.yaml",
	}

	for _, loc := range locations {
		if _, err := os.Stat(loc); err == nil {
			absPath, err := filepath.Abs(loc)
			if err == nil {
				return absPath
			}
		}
	}

	return "configs/config.yaml" // default location
}

// setDefaults sets default values if they are not provided in the config file
func setDefaults(config *Config) {
	if config.Server.Port == 0 {
		config.Server.Port = 8080
	}
	if config.Server.Host == "" {
		config.Server.Host = "0.0.0.0"
	}

	if len(config.Cors.AllowedOrigins) == 0 {
		config.Cors.AllowedOrigins = []string{"*"}
	}
	if len(config.Cors.AllowedMethods) == 0 {
		config.Cors.AllowedMethods = []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		}
	}
	if len(config.Cors.AllowedHeaders) == 0 {
		config.Cors.AllowedHeaders = []string{
			"Accept",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		}
	}

	if config.Logging.Level == "" {
		config.Logging.Level = "info"
	}
	if config.Logging.Format == "" {
		config.Logging.Format = "json"
	}

	if config.Metrics.Path == "" {
		config.Metrics.Path = "/metrics"
	}
}
