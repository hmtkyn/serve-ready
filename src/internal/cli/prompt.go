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

	// Database selection
	databaseOptions, err := requirements.ListFilesInDirectory("databases")
	if err != nil {
		return "", "", "", "", err
	}

	selectedDatabase := ""
	prompt = &survey.Select{
		Message: "Would you like to select a database? (Optional)",
		Options: append([]string{"Skip"}, databaseOptions...),
	}
	survey.AskOne(prompt, &selectedDatabase)
	if selectedDatabase == "Skip" {
		selectedDatabase = ""
	}
	fmt.Printf("Selected database: %s\n\n", selectedDatabase)

	// Cache selection
	cacheOptions, err := requirements.ListFilesInDirectory("caches")
	if err != nil {
		return "", "", "", "", err
	}

	selectedCache := ""
	prompt = &survey.Select{
		Message: "Would you like to select a cache? (Optional)",
		Options: append([]string{"Skip"}, cacheOptions...),
	}
	survey.AskOne(prompt, &selectedCache)
	if selectedCache == "Skip" {
		selectedCache = ""
	}
	fmt.Printf("Selected cache: %s\n\n", selectedCache)

	// Web server selection
	selectedWebserver := ""
	prompt = &survey.Select{
		Message: "Would you like to select a web server? (Optional)",
		Options: append([]string{"Skip"}, []string{"Nginx", "Apache"}...),
	}
	survey.AskOne(prompt, &selectedWebserver)
	if selectedWebserver == "Skip" {
		selectedWebserver = ""
	}
	fmt.Printf("Selected web server: %s\n\n", selectedWebserver)

	return selectedFramework, selectedDatabase, selectedCache, selectedWebserver, nil
}
