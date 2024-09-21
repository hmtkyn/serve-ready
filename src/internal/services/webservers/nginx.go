package services

import (
	"fmt"
	"os/exec"
)

func CheckNginxVersion(requiredVersion string) bool {
	fmt.Printf("Nginx Requirement: %s ", requiredVersion)
	nginxVersion, err := exec.Command("nginx", "-v").Output()
	if err != nil {
		fmt.Printf("%s Nginx is not installed or version mismatch.\n", redCross)
		return false
	}
	fmt.Printf("%s Nginx version is compatible: %s\n", greenCheck, string(nginxVersion))
	return true
}
