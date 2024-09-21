package requirements

import (
	"fmt"
	"serve-ready/src/internal/services/caches"
	"serve-ready/src/internal/services/databases"
	"serve-ready/src/internal/services/runtiems"
	"serve-ready/src/internal/services/webservers"
)

const (
	greenCheck = "\033[32m✔\033[0m"
	redCross   = "\033[31m✘\033[0m"
)

func CheckRequirements(framework, database, cache, webserver string) bool {
	allPassed := true

	requirements, err := LoadFrameworkRequirements(framework)
	if err != nil {
		fmt.Printf("Could not load framework requirements: %v\n", err)
		return false
	}

	fmt.Printf("\n--- Checking Requirements for %s ---\n", framework)

	// PHP Check
	if requirements.PHPVersion != "" {
		if !services.CheckPHPVersion(requirements.PHPVersion) {
			allPassed = false
		} else {
			if len(requirements.PHPExtensions) > 0 && !services.CheckPHPExtensions(requirements.PHPExtensions) {
				allPassed = false
			}
			if len(requirements.ComposerPackages) > 0 && !services.CheckComposerPackages(requirements.ComposerPackages) {
				allPassed = false
			}
		}
	}

	// Node.js Check
	if requirements.NodeVersion != "" {
		if !services.CheckNodeVersion(requirements.NodeVersion) {
			allPassed = false
		} else {
			if len(requirements.NodePackages) > 0 && !services.CheckNodePackages(requirements.NodePackages) {
				allPassed = false
			}
		}
	}

	// Python Check
	if requirements.PythonVersion != "" {
		if !services.CheckPythonVersion(requirements.PythonVersion) {
			allPassed = false
		} else {
			if len(requirements.PythonPackages) > 0 && !services.CheckPythonPackages(requirements.PythonPackages) {
				allPassed = false
			}
		}
	}

	// Database Check
	if database != "" {
		switch database {
		case "mysql":
			if !services.CheckMySQLVersion(requirements.MySQLVersion) {
				allPassed = false
			}
		case "mariadb":
			if !services.CheckMariaDBVersion(requirements.MariaDBVersion) {
				allPassed = false
			}
		case "postgresql":
			if !services.CheckPostgreSQLVersion(requirements.PostgreSQLVersion) {
				allPassed = false
			}
		case "mongodb":
			if !services.CheckMongoDBVersion(requirements.MongoDBVersion) {
				allPassed = false
			}
		default:
			fmt.Printf("%s Unsupported database: %s\n", redCross, database)
			allPassed = false
		}
	}

	// Cache Check
	if cache != "" {
		switch cache {
		case "redis":
			if !services.CheckRedisVersion(requirements.RedisVersion) {
				allPassed = false
			}
		case "firebase":
			if !services.CheckFirebaseCLI() {
				allPassed = false
			}
		default:
			fmt.Printf("%s Unsupported cache: %s\n", redCross, cache)
			allPassed = false
		}
	}

	// Web Server Check
	if webserver != "" {
		switch webserver {
		case "nginx":
			if !services.CheckNginxVersion(requirements.NginxVersion) {
				allPassed = false
			}
		case "apache":
			if !services.CheckApacheVersion(requirements.ApacheVersion) {
				allPassed = false
			}
		default:
			fmt.Printf("%s Unsupported web server: %s\n", redCross, webserver)
			allPassed = false
		}
	}

	return allPassed
}
