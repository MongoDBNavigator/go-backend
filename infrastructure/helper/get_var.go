package helper

import "os"

// Get environment variable or default value
func GetVar(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
