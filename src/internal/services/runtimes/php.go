package runtimes

import (
	"bytes"
	"fmt"
	"os/exec"
	"serve-ready/src/pkg/config"
	"strings"
)

func CheckPHP(requiredVersion string, requiredExtensions []string) bool {
	key := "PHP"
	value := "Checking configuration"

	fmt.Printf("%s: %s\n", config.Colorize(key, config.Cyan), config.Colorize(value, config.Yellow))

	var stderr bytes.Buffer
	cmd := exec.Command("php", "-v")
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil || !checkVersion(stderr.String(), requiredVersion) {
		fmt.Printf("%s: %s %s\n", config.Colorize("Error", config.Red), config.Colorize("PHP version is not compatible", config.Yellow), config.Colorize("✖", config.Red))
		return false
	}

	cmd = exec.Command("php", "-m")
	modulesOutput, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s: %s %s\n", config.Colorize("Error", config.Red), config.Colorize("Failed to get PHP modules", config.Yellow), config.Colorize("✖", config.Red))
		return false
	}

	for _, ext := range requiredExtensions {
		if !strings.Contains(string(modulesOutput), ext) {
			fmt.Printf("%s: %s %s\n", config.Colorize("Error", config.Red), config.Colorize(fmt.Sprintf("Missing PHP extension: %s", ext), config.Yellow), config.Colorize("✖", config.Red))
			return false
		}
	}

	fmt.Printf("%s: %s %s\n", config.Colorize("PHP", config.Cyan), config.Colorize("Version and extensions are compatible", config.Green), config.Colorize("✔", config.Green))
	return true
}

func checkVersion(output, requiredVersion string) bool {
	versionLine := strings.Split(output, "\n")[0]
	parts := strings.Fields(versionLine)
	if len(parts) >= 2 {
		version := parts[1]
		return version >= requiredVersion
	}
	return false
}
