package osinfo

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// Command represents a single command with a description
type Command struct {
	Description string `json:"description"`
	Command     string `json:"command"`
}

// CommandsList represents the list of commands from the JSON file
type CommandsList struct {
	Commands []Command `json:"commands"`
}

// Function to check if a command is available on the system
func checkCommandAvailability(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

// Function to run a shell command and return its output as a string
func runCommand(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

// Run executes commands from the provided JSON file
func Run(commandsFile string) {
	// Read the JSON file
	jsonFile, err := os.Open(commandsFile)
	if err != nil {
		fmt.Printf("Error opening JSON file %s: %v\n", commandsFile, err)
		return
	}
	defer jsonFile.Close()

	// Read the file's content
	byteValue, _ := io.ReadAll(jsonFile)

	// Unmarshal the JSON content into CommandsList
	var commandsList CommandsList
	err = json.Unmarshal(byteValue, &commandsList)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	// Loop through each command, check availability, and execute
	for _, cmdInfo := range commandsList.Commands {
		// Split the command to get the executable
		cmdParts := strings.Fields(cmdInfo.Command)
		executable := cmdParts[0]

		// Check if the command is available
		if !checkCommandAvailability(executable) {
			fmt.Printf("Command not available: %s\n", executable)
			continue
		}

		// Run the command and print the result
		fmt.Printf("\n---- %s ----\n", cmdInfo.Description)
		output, err := runCommand(cmdInfo.Command)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Println(output)
		}
	}
}
