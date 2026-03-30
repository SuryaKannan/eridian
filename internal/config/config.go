package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Languages      []string `json:"languages"`
	ActiveLanguage string   `json:"activeLanguage"`
}

func EridianHome() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".eridian")
}

func writeConfig(path string, c *Config) error {

	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to write eridian config: %w", err)
	}

	err = os.WriteFile(filepath.Join(path, "config.json"), jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write eridian config: %w", err)
	}

	return nil
}

func FetchConfig() *Config {

	eridianHome := EridianHome()

	content, _ := os.ReadFile(filepath.Join(eridianHome, "config.json"))

	var config Config
	json.Unmarshal(content, &config)

	return &config
}

func SetActiveLanguage(language string) error {
	cfg := FetchConfig()
	cfg.ActiveLanguage = language
	return writeConfig(EridianHome(), cfg)
}
