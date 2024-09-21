package services

import (
	"fmt"
	"os/exec"
)

func CheckMongoDBVersion(requiredVersion string) bool {
	fmt.Printf("MongoDB Requirement: %s ", requiredVersion)
	mongodbVersion, err := exec.Command("mongo", "--version").Output()
	if err != nil {
		fmt.Printf("%s MongoDB is not installed or version mismatch.\n", redCross)
		return false
	}
	fmt.Printf("%s MongoDB version is compatible: %s\n", greenCheck, string(mongodbVersion))
	return true
}
