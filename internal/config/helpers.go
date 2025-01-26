package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func getFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("issue finding user home dir: %w", err)
	}
	return filepath.Join(home, configFile), nil
}

func write(c *Config) error {
	path, err := getFilePath()
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("issue creating config file:%w", err)
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	if err = encoder.Encode(c); err != nil {
		return fmt.Errorf("issue writing data to file: %w", err)
	}

	return nil
}
