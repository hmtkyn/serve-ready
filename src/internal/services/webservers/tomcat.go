package webservers

import (
	"fmt"
	"os/exec"
	"serve-ready/src/pkg/config"
)

func CheckTomcat() bool {
	key := "Tomcat"
	value := "Checking configuration"

	fmt.Printf("%s: %s\n", config.Colorize(key, config.Cyan), config.Colorize(value, config.Yellow))

	cmd := exec.Command("catalina", "version")
	err := cmd.Run()

	if err != nil {
		fmt.Printf("%s: %s %s\n", config.Colorize("Error", config.Red), config.Colorize("Tomcat is not installed or configured.", config.Yellow), config.Colorize("✖", config.Red))
		return false
	}

	fmt.Printf("%s: %s %s\n", config.Colorize("Tomcat", config.Cyan), config.Colorize("Configured correctly", config.Green), config.Colorize("✔", config.Green))
	return true
}
