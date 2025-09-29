package envy

import (
	"fmt"
	"os"
	"time"
)

// getEnv retrieves the value of the environment variable named by the key.
// If the variable is not present in the environment, then the provided default value is returned.
// It first checks for the variable using os.LookupEnv, and if not found, it falls back to os.Getenv.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	if nextValue := os.Getenv(key); nextValue != "" {
		return nextValue
	}

	return defaultValue

}

// Get retrieves the value of the environment variable named by the key.
// If the variable is not present in the environment, then the provided default value is returned.
func Get(key, defaultValue string) string {
	return getEnv(key, defaultValue)
}

// GetAsBool retrieves the value of the environment variable named by the key and converts it to a boolean.
// If the variable is not present in the environment, then the provided default value is returned.
func GetAsBool(key string, defaultValue bool) bool {
	valStr := getEnv(key, "")
	
	if valStr == "" {
		return defaultValue
	}

	if valStr == "true" || valStr == "1" {
		return true
	}

	return false
}

// GetAsInt retrieves the value of the environment variable named by the key and converts it to an integer.
// If the variable is not present in the environment, then the provided default value is returned.
func GetAsInt(key string, defaultValue int) int {
	valStr := getEnv(key, "")
	if valStr == "" {
		return defaultValue
	}

	var intValue int
	_, err := fmt.Sscanf(valStr, "%d", &intValue)
	if err != nil {
		return defaultValue
	}

	return intValue
}

// GetAsFloat retrieves the value of the environment variable named by the key and converts it to a float.
// If the variable is not present in the environment, then the provided default value is returned.
func GetAsFloat(key string, defaultValue float64) float64 {
	valStr := getEnv(key, "")
	if valStr == "" {
		return defaultValue
	}

	var floatValue float64
	_, err := fmt.Sscanf(valStr, "%f", &floatValue)
	if err != nil {
		return defaultValue
	}

	return floatValue
}

// GetAsDuration retrieves the value of the environment variable named by the key and converts it to a time.Duration.
// If the variable is not present in the environment, then the provided default value is returned.
func GetAsDuration(key string, defaultValue string) time.Duration {
	valStr := getEnv(key, "")
	if valStr == "" {
		valStr = defaultValue
	}

	durationValue, _ := time.ParseDuration(valStr)
	return durationValue
}
