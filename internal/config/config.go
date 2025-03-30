package config

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

type Config struct {
	BasePath    string `json:"base_path"`
	DefaultLang string `json:"default_lang"`
	GitUser     string `json:"git_user"`
	GitEmail    string `json:"git_email"`
}

func NewConfig() (*Config, error) {
	currentUser, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("failed to get current user: %w", err)
	}

	configPath := filepath.Join(currentUser.HomeDir, ".qs", "config.json")
	cfg := &Config{
		BasePath: filepath.Join(currentUser.HomeDir, "Documents"),
	}

	if _, err := os.Stat(configPath); err == nil {
		data, err := os.ReadFile(configPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
		if err := json.Unmarshal(data, cfg); err != nil {
			return nil, fmt.Errorf("failed to parse config file: %w", err)
		}
	}

	return cfg, nil
}

func (c *Config) Save() error {
	currentUser, err := user.Current()
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}

	configDir := filepath.Join(currentUser.HomeDir, ".qs")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	configPath := filepath.Join(configDir, "config.json")
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

func (c *Config) CheckDependencies() error {
	deps := []string{"git", "code"}
	for _, dep := range deps {
		if _, err := os.Stat(dep); err != nil {
			return fmt.Errorf("dependency %s not found: %w", dep, err)
		}
	}
	return nil
}
