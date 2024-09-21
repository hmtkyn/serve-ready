package caches

import (
	"fmt"
	"os/exec"
	"serve-ready/src/pkg/config"
)

func CheckRedisCache() bool {
	key := "Redis Cache"
	value := "Checking configuration"

	fmt.Printf("%s: %s\n", config.Colorize(key, config.Cyan), config.Colorize(value, config.Yellow))

	cmd := exec.Command("redis-server", "--version")
	err := cmd.Run()

	if err != nil {
		fmt.Printf("%s: %s %s\n", config.Colorize("Error", config.Red), config.Colorize("Redis is not configured.", config.Yellow), config.Colorize("✖", config.Red))
		return false
	}

	fmt.Printf("%s: %s %s\n", config.Colorize("Redis Cache", config.Cyan), config.Colorize("Configured correctly", config.Green), config.Colorize("✔", config.Green))
	return true
}
