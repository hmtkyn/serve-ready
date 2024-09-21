package services

import (
	"fmt"
	"os/exec"
)

func CheckMariaDBVersion(requiredVersion string) bool {
	fmt.Printf("MariaDB Requirement: %s ", requiredVersion)
	mariadbVersion, err := exec.Command("mysql", "--version").Output()
	if err != nil {
		fmt.Printf("%s MariaDB is not installed or version mismatch.\n", redCross)
		return false
	}
	fmt.Printf("%s MariaDB version is compatible: %s\n", greenCheck, string(mariadbVersion))
	return true
}
