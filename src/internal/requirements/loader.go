package requirements

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

// FrameworkRequirements defines the structure for a framework's requirements
type FrameworkRequirements struct {
	PHPVersion    string   `yaml:"php_version,omitempty"`
	NodeVersion   string   `yaml:"node_version,omitempty"`
	PythonVersion string   `yaml:"python_version,omitempty"`
	Extensions    []string `yaml:"extensions,omitempty"`
	Packages      []string `yaml:"packages,omitempty"`
	NPMRequired   bool     `yaml:"npm_required,omitempty"`
}

// ListFilesInDirectory lists the .yml files in a given directory
func ListFilesInDirectory(dir string) ([]string, error) {
	// Get the current working directory
	basePath, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("could not determine working directory: %w", err)
	}

	// Combine the base path with the requested directory
	fullPath := filepath.Join(basePath, "configs", dir) // Adjusted to point to 'configs/'

	// Read the directory
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

// LoadFrameworkRequirements loads the requirements from a .yml file
func LoadFrameworkRequirements(framework string) (*FrameworkRequirements, error) {
	// Get the current working directory
	basePath, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("could not determine working directory: %w", err)
	}

	// Combine base path with the frameworks directory
	filePath := filepath.Join(basePath, "configs", "frameworks", fmt.Sprintf("%s.yml", framework))

	// Read the file
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
