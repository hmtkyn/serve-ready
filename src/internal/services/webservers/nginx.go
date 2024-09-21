package webservers

import (
	"bytes"
	"fmt"
	"os/exec"
	"serve-ready/src/pkg/config"
)

func CheckNginx() bool {
	key := "Nginx"
	value := "Checking configuration"

	fmt.Printf("%s: %s\n", config.Colorize(key, config.Cyan), config.Colorize(value, config.Yellow))

	cmd := exec.Command("nginx", "-v")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		fmt.Printf("%s: %s %s\n", config.Colorize("Error", config.Red), config.Colorize("Nginx is not installed or configured.", config.Yellow), config.Colorize("✖", config.Red))
		return false
	}

	fmt.Printf("%s: %s %s\n", config.Colorize("Nginx", config.Cyan), config.Colorize(stderr.String(), config.Green), config.Colorize("✔", config.Green))
	return true
}
