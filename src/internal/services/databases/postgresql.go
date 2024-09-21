package services

import (
	"fmt"
	"os/exec"
)

func CheckPostgreSQLVersion(requiredVersion string) bool {
	fmt.Printf("PostgreSQL Requirement: %s ", requiredVersion)
	postgresqlVersion, err := exec.Command("psql", "--version").Output()
	if err != nil {
		fmt.Printf("%s PostgreSQL is not installed or version mismatch.\n", redCross)
		return false
	}
	fmt.Printf("%s PostgreSQL version is compatible: %s\n", greenCheck, string(postgresqlVersion))
	return true
}
