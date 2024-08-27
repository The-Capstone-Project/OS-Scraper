package main

import (
	"fmt"
	"os/exec"
	"strings"
)

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

func main() {
	// List of commands to check and execute
	commands := map[string]string{
		"Distribution Info":       "lsb_release -a || cat /etc/os-release",
		"Kernel Version":          "uname -r",
		"Full Kernel Information": "uname -a",
		"GRUB Settings":           "cat /etc/default/grub",
		"CPU Information":         "lscpu",
		"Memory Information":      "free -h",
		"Disk Usage":              "df -h",
		"Network Interfaces":      "ip a",
		"Routing Table":           "ip route",
		"System Information":      "hostnamectl",
		"File Systems":            "mount | column -t",
		"Installed Packages":      "dpkg -l",
		"Running Processes":       "ps aux",
		"Active Services":         "systemctl list-units --type=service --state=running",
		"Logged In Users":         "who",
		"Environment Variables":   "printenv",
	}

	// Loop through each command, check availability, and execute
	for desc, cmd := range commands {
		// Split the command to get the executable (e.g., "lsb_release")
		cmdParts := strings.Fields(cmd)
		executable := cmdParts[0]

		// Check if the command is available
		if !checkCommandAvailability(executable) {
			fmt.Printf("Command not available: %s\n", executable)
			continue
		}

		// Run the command and print the result
		fmt.Printf("\n---- %s ----\n", desc)
		output, err := runCommand(cmd)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Println(output)
		}
	}
}
