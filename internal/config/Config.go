package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFile string = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("issue finding user home dir: %w", err)
	}
	return filepath.Join(home, configFile), nil
}

func Read() (*Config, error) {
	path, err := getFilePath()
	if err != nil {
		return nil, err
	}

	f, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("issue reading json file: %w", err)
	}

	config := &Config{}
	json.Unmarshal(f, config)

	return config, nil
}

func write(c *Config) error {
	path, err := getFilePath()
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	defer f.Close()

	if err != nil {
		return fmt.Errorf("issue creating config file:%w", err)
	}

	data, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("issue marshaling data: %w", err)
	}

	if _, err := f.Write(data); err != nil {
		return fmt.Errorf("issu writing data to file: %w", err)
	}

	return nil
}

func (c *Config) SetUser(name string) {
	c.CurrentUserName = name
	write(c)
}
