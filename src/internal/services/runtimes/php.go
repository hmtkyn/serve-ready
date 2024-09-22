package runtimes

import (
	"fmt"
	"os/exec"
	"regexp"
	"serve-ready/src/pkg/config"
	"strconv"
	"strings"
)

// CheckPHP checks if the correct PHP version and extensions are installed
func CheckPHP(requiredVersion string, requiredExtensions []string) bool {
	fmt.Println("PHP: Checking configuration")

	cmd := exec.Command("php", "--version")
	output, err := cmd.Output()

	if err != nil {
		fmt.Printf("%s PHP is not installed or configured properly.\n", config.Colorize("Error", config.Red))
		return false
	}

	// Extract PHP version
	versionRegex := regexp.MustCompile(`PHP\s([0-9]+)\.([0-9]+)\.([0-9]+)`)
	matches := versionRegex.FindStringSubmatch(string(output))

	if len(matches) < 4 {
		fmt.Println(config.Colorize("Error: Could not parse PHP version.", config.Red))
		return false
	}

	major, _ := strconv.Atoi(matches[1])
	minor, _ := strconv.Atoi(matches[2])

	// Parse required version (>=8.1)
	requiredMajor, requiredMinor := parsePHPVersion(requiredVersion)
	if major < requiredMajor || (major == requiredMajor && minor < requiredMinor) {
		fmt.Printf("%s Required PHP version is %s, but installed version is %d.%d ✖\n", config.Colorize("Error", config.Red), requiredVersion, major, minor)
		return false
	}

	// Check required PHP extensions
	for _, ext := range requiredExtensions {
		if !checkPHPExtension(ext) {
			fmt.Printf("%s PHP extension %s is missing ✖\n", config.Colorize("Error", config.Red), ext)
			return false
		}
	}

	fmt.Println(config.Colorize("PHP version and extensions are compatible ✔", config.Green))
	return true
}

// parsePHPVersion parses the required PHP version string (e.g., ">=8.1")
func parsePHPVersion(version string) (int, int) {
	version = strings.TrimPrefix(version, ">=")
	parts := strings.Split(version, ".")
	major, _ := strconv.Atoi(parts[0])
	minor, _ := strconv.Atoi(parts[1])
	return major, minor
}

// checkPHPExtension checks if a PHP extension is installed
func checkPHPExtension(extension string) bool {
	cmd := exec.Command("php", "-m")
	output, _ := cmd.Output()

	extensions := strings.Split(string(output), "\n")
	for _, ext := range extensions {
		if strings.TrimSpace(ext) == extension {
			return true
		}
	}

	return false
}
