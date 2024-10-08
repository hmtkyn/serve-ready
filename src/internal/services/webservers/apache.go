package webservers

import (
	"fmt"
	"os/exec"
	"serve-ready/src/pkg/config"
)

func CheckApache() bool {
	key := "Apache"
	value := "Checking configuration"

	fmt.Printf("%s: %s\n", config.Colorize(key, config.Cyan), config.Colorize(value, config.Yellow))

	cmd := exec.Command("apachectl", "-v")
	err := cmd.Run()

	if err != nil {
		fmt.Printf("%s: %s %s\n", config.Colorize("Error", config.Red), config.Colorize("Apache is not installed or configured.", config.Yellow), config.Colorize("✖", config.Red))
		return false
	}

	fmt.Printf("%s: %s %s\n", config.Colorize("Apache", config.Cyan), config.Colorize("Configured correctly", config.Green), config.Colorize("✔", config.Green))
	return true
}
