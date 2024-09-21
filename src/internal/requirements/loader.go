package requirements

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type FrameworkRequirements struct {
	PHPVersion         string   `yaml:"php_version,omitempty"`
	RequiredExtensions []string `yaml:"required_extensions,omitempty"`
}

func ListFilesInDirectory(dir string) ([]string, error) {
	basePath, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("could not determine working directory: %w", err)
	}

	var fullPath string
	if dir == "frameworks" {
		// Use the correct path for frameworks
		fullPath = filepath.Join(basePath, "src", "internal", "frameworks")
	} else {
		// Use the correct path for services (databases, caches, etc.)
		fullPath = filepath.Join(basePath, "src", "internal", "services", dir)
	}

	files, err := os.ReadDir(fullPath)
	if err != nil {
		return nil, fmt.Errorf("could not read directory %s: %w", fullPath, err)
	}

	var options []string
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".yml" {
			option := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
			options = append(options, option)
		}
	}
	return options, nil
}

func LoadFrameworkRequirements(framework string) (*FrameworkRequirements, error) {
	basePath, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("could not determine working directory: %w", err)
	}

	filePath := filepath.Join(basePath, "src", "internal", "frameworks", fmt.Sprintf("%s.yml", framework))

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not read file %s: %w", filePath, err)
	}

	var reqs FrameworkRequirements
	err = yaml.Unmarshal(data, &reqs)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal YAML: %w", err)
	}

	return &reqs, nil
}
