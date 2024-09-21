package services

import (
	"fmt"
	"os/exec"
	"strings"
)

// PHPVersionCheck checks if the PHP version matches the required version
func PHPVersionCheck(requiredVersion string) (bool, string) {
	if requiredVersion == "" {
		return true, ""
	}

	phpVersion, err := exec.Command("php", "-v").Output()
	if err != nil {
		return false, "PHP is not installed or not found in PATH"
	}

	installedVersion := parsePHPVersion(string(phpVersion))
	if !strings.Contains(installedVersion, requiredVersion) {
		return false, fmt.Sprintf("PHP version mismatch. Installed: %s, Required: %s", installedVersion, requiredVersion)
	}

	return true, installedVersion
}

// parsePHPVersion extracts the PHP version from the command output
func parsePHPVersion(output string) string {
	lines := strings.Split(output, "\n")
	if len(lines) > 0 {
		parts := strings.Split(lines[0], " ")
		for _, part := range parts {
			if strings.Count(part, ".") == 2 {
				return part
			}
		}
	}
	return ""
}
