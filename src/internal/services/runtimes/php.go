package services

import (
	"fmt"
	"os/exec"
	"strings"
)

func CheckPHPVersion(requiredVersion string) bool {
	fmt.Printf("PHP Requirement: %s ", requiredVersion)
	phpVersion, err := CheckVersion("php", "-v")
	if err != nil || !strings.Contains(phpVersion, requiredVersion) {
		fmt.Printf("%s PHP is not installed or version mismatch.\n", redCross)
		return false
	}
	fmt.Printf("%s PHP version is compatible: %s\n", greenCheck, phpVersion)
	return true
}

func CheckPHPExtensions(requiredExtensions []string) bool {
	allPassed := true
	for _, ext := range requiredExtensions {
		_, err := exec.Command("php", "-m").Output()
		if err != nil || !strings.Contains(string(err.Error()), ext) {
			fmt.Printf("%s PHP Extension '%s' is missing.\n", redCross, ext)
			allPassed = false
		} else {
			fmt.Printf("%s PHP Extension '%s' is installed.\n", greenCheck, ext)
		}
	}
	return allPassed
}

func CheckComposerPackages(requiredPackages []string) bool {
	allPassed := true
	for _, pkg := range requiredPackages {
		_, err := exec.Command("composer", "show", pkg).Output()
		if err != nil {
			fmt.Printf("%s Composer Package '%s' is missing.\n", redCross, pkg)
			allPassed = false
		} else {
			fmt.Printf("%s Composer Package '%s' is installed.\n", greenCheck, pkg)
		}
	}
	return allPassed
}
