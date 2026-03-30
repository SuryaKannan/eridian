package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func checkEridianExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("error checking directory: %w", err)
	}
	return true, nil
}

func initialiseEridianHome(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return fmt.Errorf("failed to create eridian dir in home: %w", err)
	}

	emptyConfig := Config{
		ActiveLanguage: "",
		Languages:      []string{},
	}

	return writeConfig(path, &emptyConfig)

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

func syncCurrentLanguage(path string) error {

	content, err := os.ReadFile(filepath.Join(path, "config.json"))
	if err != nil {
		return fmt.Errorf("error reading eridian config: %w", err)
	}

	var config Config
	if err := json.Unmarshal(content, &config); err != nil {
		return fmt.Errorf("error parsing eridian config: %w", err)
	}

	activeLanguage := config.ActiveLanguage

	entries, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("failed to read eridian directory: %w", err)
	}

	languages := []string{}
	resolvedLanguage := ""

	for _, entry := range entries {

		if strings.HasSuffix(entry.Name(), ".db") {
			language := strings.TrimSuffix(entry.Name(), ".db")
			languages = append(languages, language)

			if activeLanguage == language {
				resolvedLanguage = activeLanguage
			}
		}
	}

	if resolvedLanguage == config.ActiveLanguage &&
		slices.Equal(languages, config.Languages) {
		return nil
	}

	updatedConfig := Config{
		ActiveLanguage: resolvedLanguage,
		Languages:      languages,
	}

	return writeConfig(path, &updatedConfig)

}

func SyncConfig() error {
	home, err := os.UserHomeDir()

	if err != nil {
		return fmt.Errorf("could not find home directory: %w", err)
	}

	eridianHome := filepath.Join(home, ".eridian")
	dirExists, err := checkEridianExists(eridianHome)

	if err != nil {
		return fmt.Errorf("error checking eridian home: %w", err)
	}

	if !dirExists {
		return initialiseEridianHome(eridianHome)
	}

	return syncCurrentLanguage(eridianHome)
}
