package requirements

import (
	"fmt"
	"strings"
)

// CheckPython checks if the Python version matches the required version
func CheckPython(requiredVersion string) bool {
	if requiredVersion == "" {
		return true
	}
	fmt.Printf("Python Requirement: %s ", requiredVersion)
	pythonVersion, err := CheckVersion("python3", "--version")
	if err != nil || !strings.Contains(pythonVersion, requiredVersion) {
		fmt.Printf("%s Python is not installed or version mismatch.\n", redCross)
		return false
	}
	fmt.Printf("%s Python version is compatible: %s\n", greenCheck, pythonVersion)
	return true
}
