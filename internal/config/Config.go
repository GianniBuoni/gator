package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Read() (*Config, error) {
	path, err := getFilePath()
	if err != nil {
		return nil, err
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("issue opening file: %w", err)
	}
	defer f.Close()

	config := &Config{}
	decoder := json.NewDecoder(f)
	if err = decoder.Decode(config); err != nil {
		return nil, fmt.Errorf("issue decoding file contents: %w", err)
	}

	return config, nil
}

func (c *Config) SetUser(name string) {
	c.CurrentUserName = name
	write(c)
}
