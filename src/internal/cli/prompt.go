package cli

import (
	"fmt"
	"serve-ready/src/internal/requirements"

	"github.com/AlecAivazis/survey/v2"
)

// GetSelections handles prompting the user for selections
func GetSelections() (string, string, string, string, error) {
	// Framework selection
	frameworkOptions, err := requirements.ListFilesInDirectory("frameworks")
	if err != nil {
		return "", "", "", "", err
	}

	selectedFramework := ""
	prompt := &survey.Select{
		Message: "Select a framework:",
		Options: frameworkOptions,
	}
	survey.AskOne(prompt, &selectedFramework)
	fmt.Printf("Selected framework: %s\n\n", selectedFramework)

	// Hardcoded database options
	databaseOptions := []string{"MySQL", "PostgreSQL", "Skip"}
	selectedDatabase := ""
	prompt = &survey.Select{
		Message: "Would you like to select a database? (Optional)",
		Options: databaseOptions,
	}
	survey.AskOne(prompt, &selectedDatabase)
	if selectedDatabase == "Skip" {
		selectedDatabase = ""
	}
	fmt.Printf("Selected database: %s\n\n", selectedDatabase)

	// Hardcoded cache options
	cacheOptions := []string{"Redis", "Memcached", "Skip"}
	selectedCache := ""
	prompt = &survey.Select{
		Message: "Would you like to select a cache? (Optional)",
		Options: cacheOptions,
	}
	survey.AskOne(prompt, &selectedCache)
	if selectedCache == "Skip" {
		selectedCache = ""
	}
	fmt.Printf("Selected cache: %s\n\n", selectedCache)

	// Hardcoded web server options
	webserverOptions := []string{"Nginx", "Apache", "Skip"}
	selectedWebserver := ""
	prompt = &survey.Select{
		Message: "Would you like to select a web server? (Optional)",
		Options: webserverOptions,
	}
	survey.AskOne(prompt, &selectedWebserver)
	if selectedWebserver == "Skip" {
		selectedWebserver = ""
	}
	fmt.Printf("Selected web server: %s\n\n", selectedWebserver)

	return selectedFramework, selectedDatabase, selectedCache, selectedWebserver, nil
}
