package env_info

import (
	"fmt"
	"os"
	"strings"
)

func Run() {
	// Fetch all environment variables
	envVars := os.Environ()

	// Define a list of known private API key environment variables
	privateKeys := []string{
		"API_KEY",
		"SECRET_KEY",
		"PRIVATE_KEY",
		"PASSWORD",
	}

	// Convert private keys to lowercase for case-insensitive comparison
	privateKeySet := make(map[string]struct{})
	for _, key := range privateKeys {
		privateKeySet[strings.ToLower(key)] = struct{}{}
	}

	// Iterate over environment variables
	fmt.Println("Environment Information:")
	for _, env := range envVars {
		// Split key and value
		parts := strings.SplitN(env, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := parts[0]
		value := parts[1]

		// Check if the key is in the private keys list
		if _, found := privateKeySet[strings.ToLower(key)]; found {
			fmt.Printf("%s: [REDACTED]\n", key)
		} else {
			fmt.Printf("%s: %s\n", key, value)
		}
	}
}
