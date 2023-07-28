package env

import (
	"errors"
	"os"
)

// MustGetEnv takes a key string and returns an error if it is not found, and a string value if otherwise
func MustGetEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", errors.New("can't find environment variable: " + key)
	}
	return value, nil
}

// GetEnv takes a key string and returns a string
// if the environment isn't found, then the default value provided is used.
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
