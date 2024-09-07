package main

import (
	"bufio"
	"fmt"
	"os"

	"strings"
)

func main() {
	// Ask the user whether to fetch history or not
	fmt.Print("Do you want to fetch the history? (yes/no): ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Trim any whitespace and convert input to lowercase
	input = strings.TrimSpace(strings.ToLower(input))

	if input == "no" {
		fmt.Println("User chose not to execute the history command.")
		return
	} else if input == "yes" || input == "y" {

		// Print the output of 'history' command
		fmt.Println("Output of 'history' command:")

	} else {
		fmt.Println("Invalid input, exiting program.")
		return
	}

	// Read the .zsh_history file and save its content to a variable
	zshHistoryFile := os.Getenv("HOME") + "/.zsh_history"
	zshHistoryContent, err := os.ReadFile(zshHistoryFile)
	if err != nil {
		fmt.Println("Error reading .zsh_history file:", err)
		return
	}
	zshHistoryData := string(zshHistoryContent)

	// Print the content of .zsh_history file
	fmt.Println("Content of .zsh_history file:")
	fmt.Println(zshHistoryData)
}
