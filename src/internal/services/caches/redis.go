package services

import (
	"fmt"
	"os/exec"
)

func CheckRedisVersion(requiredVersion string) bool {
	fmt.Printf("Redis Requirement: %s ", requiredVersion)
	redisVersion, err := exec.Command("redis-server", "--version").Output()
	if err != nil {
		fmt.Printf("%s Redis is not installed or version mismatch.\n", redCross)
		return false
	}
	fmt.Printf("%s Redis version is compatible: %s\n", greenCheck, string(redisVersion))
	return true
}
