package services

import (
	"fmt"
	"os/exec"
	"strings"
)

const (
	greenCheck = "\033[32m✔\033[0m"
	redCross   = "\033[31m✘\033[0m"
)

func CheckPythonVersion(requiredVersion string) bool {
	fmt.Printf("Python Requirement: %s ", requiredVersion)
	pythonVersion, err := CheckVersion("python3", "--version")
	if err != nil || !strings.Contains(pythonVersion, requiredVersion) {
		fmt.Printf("%s Python is not installed or version mismatch.\n", redCross)
		return false
	}
	fmt.Printf("%s Python version is compatible: %s\n", greenCheck, pythonVersion)
	return true
}

func CheckPythonPackages(requiredPackages []string) bool {
	allPassed := true
	for _, pkg := range requiredPackages {
		_, err := exec.Command("pip3", "show", pkg).Output()
		if err != nil {
			fmt.Printf("%s Package '%s' is missing.\n", redCross, pkg)
			allPassed = false
		} else {
			fmt.Printf("%s Package '%s' is installed.\n", greenCheck, pkg)
		}
	}
	return allPassed
}
