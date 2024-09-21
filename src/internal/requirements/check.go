package requirements

import (
	"fmt"
	"serve-ready/src/internal/services/caches"
	"serve-ready/src/internal/services/runtimes"
	"serve-ready/src/internal/services/webservers"
)

func CheckRequirements(framework, webserver, cache string) bool {
	allPassed := true

	requirements, err := LoadFrameworkRequirements(framework)
	if err != nil {
		fmt.Printf("Could not load framework requirements: %v\n", err)
		return false
	}

	if framework == "laravel" {
		if !runtimes.CheckPHP(requirements.PHPVersion, requirements.RequiredExtensions) {
			allPassed = false
		}
	}

	if webserver != "" {
		if webserver == "nginx" && !webservers.CheckNginx() {
			allPassed = false
		}
		if webserver == "apache" && !webservers.CheckApache() {
			allPassed = false
		}
	}

	if cache != "" {
		if cache == "redis" && !caches.CheckRedisCache() {
			allPassed = false
		}
	}

	return allPassed
}
