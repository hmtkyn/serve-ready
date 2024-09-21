package services

import (
	"fmt"
	"os/exec"
	"strings"
)

func CheckNodeVersion(requiredVersion string) bool {
	fmt.Printf("Node.js Requirement: %s ", requiredVersion)
	nodeVersion, err := CheckVersion("node", "-v")
	if err != nil || !strings.Contains(nodeVersion, requiredVersion) {
		fmt.Printf("%s Node.js is not installed or version mismatch.\n", redCross)
		return false
	}
	fmt.Printf("%s Node.js version is compatible: %s\n", greenCheck, nodeVersion)
	return true
}

func CheckNodePackages(requiredPackages []string) bool {
	allPassed := true
	for _, pkg := range requiredPackages {
		_, err := exec.Command("npm", "list", pkg).Output()
		if err != nil {
			fmt.Printf("%s Node Package '%s' is missing.\n", redCross, pkg)
			allPassed = false
		} else {
			fmt.Printf("%s Node Package '%s' is installed.\n", greenCheck, pkg)
		}
	}
	return allPassed
}
