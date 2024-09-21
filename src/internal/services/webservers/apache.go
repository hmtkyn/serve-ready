package services

import (
	"fmt"
	"os/exec"
)

func CheckApacheVersion(requiredVersion string) bool {
	fmt.Printf("Apache Requirement: %s ", requiredVersion)
	apacheVersion, err := exec.Command("apache2", "-v").Output()
	if err != nil {
		fmt.Printf("%s Apache is not installed or version mismatch.\n", redCross)
		return false
	}
	fmt.Printf("%s Apache version is compatible: %s\n", greenCheck, string(apacheVersion))
	return true
}
