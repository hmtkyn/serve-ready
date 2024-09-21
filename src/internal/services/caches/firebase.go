package services

import (
	"fmt"
	"os/exec"
)

func CheckFirebaseCLI() bool {
	_, err := exec.Command("firebase", "--version").Output()
	if err != nil {
		fmt.Printf("%s Firebase CLI is not installed.\n", redCross)
		return false
	}
	fmt.Printf("%s Firebase CLI is installed.\n", greenCheck)
	return true
}
