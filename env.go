package env

import (
	"fmt"
	"os"
)

// GetEnv gets an environment variable from a string and returns an error if the env variable is not set
func GetEnv(key string) (string, error) {
	variable := os.Getenv(key)
	if variable == "" {
		return "", fmt.Errorf("environment variable %s could not be found", key)
	}

	return variable, nil
}

// GetEnvOrDefault gets an environment variable from a string or return a default value
func GetEnvOrDefault(key, defaultValue string) string {
	variable := os.Getenv(key)
	if variable == "" {
		return defaultValue
	}

	return variable
}
