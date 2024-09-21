package requirements

import (
	"fmt"
	"os/exec"
	"strings"
)

// Color definitions
const (
	greenCheck = "\033[32m✔\033[0m" // Green check mark
	redCross   = "\033[31m✘\033[0m" // Red cross mark
	resetColor = "\033[0m"
)

// CheckVersion checks the version of a command
func CheckVersion(command string, args ...string) (string, error) {
	out, err := exec.Command(command, args...).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

// CheckRequirements checks the selected framework, database, and cache
func CheckRequirements(framework, database, cache, webserver string) bool {
	allPassed := true

	requirements, err := LoadFrameworkRequirements(framework)
	if err != nil {
		fmt.Printf("Could not load framework requirements: %v\n", err)
		return false
	}

	fmt.Printf("\n--- Checking Requirements for %s ---\n", framework)

	// PHP Check
	if requirements.PHPVersion != "" {
		fmt.Printf("PHP Requirement: %s ", requirements.PHPVersion)
		phpVersion, err := CheckVersion("php", "-v")
		if err != nil || !strings.Contains(phpVersion, requirements.PHPVersion) {
			fmt.Printf("%s PHP is not installed or version mismatch.\n", redCross)
			allPassed = false
		} else {
			fmt.Printf("%s PHP version is compatible: %s\n", greenCheck, phpVersion)
		}
	}

	// Node.js Check
	if requirements.NodeVersion != "" {
		fmt.Printf("Node.js Requirement: %s ", requirements.NodeVersion)
		nodeVersion, err := CheckVersion("node", "-v")
		if err != nil || !strings.Contains(nodeVersion, requirements.NodeVersion) {
			fmt.Printf("%s Node.js is not installed or version mismatch.\n", redCross)
			allPassed = false
		} else {
			fmt.Printf("%s Node.js version is compatible: %s\n", greenCheck, nodeVersion)
		}
	}

	// Python Check
	if requirements.PythonVersion != "" {
		fmt.Printf("Python Requirement: %s ", requirements.PythonVersion)
		pythonVersion, err := CheckVersion("python3", "--version")
		if err != nil || !strings.Contains(pythonVersion, requirements.PythonVersion) {
			fmt.Printf("%s Python is not installed or version mismatch.\n", redCross)
			allPassed = false
		} else {
			fmt.Printf("%s Python version is compatible: %s\n", greenCheck, pythonVersion)
		}
	}

	// Optional checks
	if database != "" {
		fmt.Printf("Database: %s %s\n", database, greenCheck)
	}

	if cache != "" {
		fmt.Printf("Cache: %s %s\n", cache, greenCheck)
	}

	if webserver != "" {
		fmt.Printf("Web Server: %s %s\n", webserver, greenCheck)
	}

	return allPassed
}
