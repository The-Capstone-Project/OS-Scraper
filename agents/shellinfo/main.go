package shellinfo

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Run() {
	// Fetch the shell environment variable
	shell := os.Getenv("SHELL")
	if shell == "" {
		log.Fatalf("Failed to determine the shell. SHELL environment variable is not set.")
	}
	fmt.Printf("Current shell: %s\n", shell)

	// Execute the whoami command
	cmd := exec.Command("whoami")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed to execute command: %s", err)
	}

	// Save the output to a variable
	userName := string(output)
	userName = strings.TrimSpace(userName) // Remove any trailing newline characters

	// Check if zsh exists in the shell path
	if strings.Contains(shell, "zsh") {
		// Construct the file path using the username
		filePath := fmt.Sprintf("/home/%s/.zshrc", userName)

		// Read the file using os.ReadFile
		content, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Failed to read file: %s", err)
		}

		// Print the file content
		fmt.Println(string(content))
	}

	// Check if bash exists in the shell path
	if strings.Contains(shell, "bash") {
		// Construct the file path using the username
		filePath := fmt.Sprintf("/home/%s/.bashrc", userName)

		// Read the file using os.ReadFile
		content, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Failed to read file: %s", err)
		}

		// Print the file content
		fmt.Println(string(content))
	}

	// Check if fish exists in the shell path
	if strings.Contains(shell, "fish") {

		// Construct the file path using the username
		filePath := fmt.Sprintf("/home/%s//.config/fish/config.fish", userName)

		// Read the file using os.ReadFile
		content, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Failed to read file: %s", err)
		}

		// Print the file content
		fmt.Println(string(content))
	}
}
