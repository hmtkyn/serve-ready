package services

import (
	"fmt"
	"os/exec"
	"strings"
)

// NodeVersionCheck checks if the Node.js version matches the required version
func NodeVersionCheck(requiredVersion string) (bool, string) {
	if requiredVersion == "" {
		return true, ""
	}

	nodeVersion, err := exec.Command("node", "-v").Output()
	if err != nil {
		return false, "Node.js is not installed or not found in PATH"
	}

	installedVersion := strings.TrimSpace(string(nodeVersion))
	if !strings.Contains(installedVersion, requiredVersion) {
		return false, fmt.Sprintf("Node.js version mismatch. Installed: %s, Required: %s", installedVersion, requiredVersion)
	}

	return true, installedVersion
}
