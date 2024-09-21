package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

// LoadConfig reads and parses the YAML configuration file
func LoadConfig(filePath string, out interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open config file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("could not read config file: %w", err)
	}

	if err := yaml.Unmarshal(data, out); err != nil {
		return fmt.Errorf("could not unmarshal YAML: %w", err)
	}

	return nil
}
