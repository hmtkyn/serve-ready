package services

import (
	"fmt"
	"os/exec"
)

func CheckMySQLVersion(requiredVersion string) bool {
	fmt.Printf("MySQL Requirement: %s ", requiredVersion)
	mysqlVersion, err := exec.Command("mysql", "--version").Output()
	if err != nil {
		fmt.Printf("%s MySQL is not installed or version mismatch.\n", redCross)
		return false
	}
	fmt.Printf("%s MySQL version is compatible: %s\n", greenCheck, string(mysqlVersion))
	return true
}
