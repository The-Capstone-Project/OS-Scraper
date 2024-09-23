package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func runCommand(cmd string, args ...string) string {
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return fmt.Sprintf("Error running command: %s\n", err)
	}
	return string(out)
}

func getActiveServices() string {
	//To Just Run The Command
	//return runCommand("systemctl", "list-units", "--type=service", "--state=running")

	//To get just the names of the services
	output := runCommand("systemctl", "list-units", "--type=service", "--state=running")
	var serviceNames []string
	scanner := bufio.NewScanner(strings.NewReader(output))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			parts := strings.Fields(line)
			if len(parts) > 0 {
				serviceNames = append(serviceNames, parts[0])
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Sprintf("Error scanning output: %s\n", err)
	}

	fmt.Printf("Total Number Of Services: %d\n", (len(serviceNames) - 5)) //Displays Total Number Of Active Services
	return strings.Join(serviceNames[1:len(serviceNames)-4], "\n")        //Returns Active Service Names Only

}

func getServiceStartupSettings() string {
	//To Just Run The Command
	//return runCommand("systemctl", "list-unit-files", "--type=service")

	//To get just the service names and their startup status ommiting the PRESET column
	output := runCommand("systemctl", "list-unit-files", "--type=service")
	var serviceStatuses []string
	scanner := bufio.NewScanner(strings.NewReader(output))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			//The service name is the first column, and the status is the second
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				serviceStatuses = append(serviceStatuses, fmt.Sprintf("%s: %s", parts[0], parts[1]))
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Sprintf("Error scanning output: %s\n", err)
	}

	fmt.Printf("Total Number Of Services: %d\n", (len(serviceStatuses) - 2)) //Displays Total Number Of Services
	return strings.Join(serviceStatuses[1:len(serviceStatuses)-1], "\n")     //Returns services with startup settings

}

func readDaemonConf(path string) string {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Sprintf("Error reading file: %s\n", err)
	}
	defer file.Close()

	var result []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// Ignore empty lines and comments (To be changed if to include comments and Empty Lines)
		if line != "" && !strings.HasPrefix(line, "#") {
			result = append(result, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Sprintf("Error scanning file: %s\n", err)
	}
	return strings.Join(result, "\n")

}

// Main fn for printing Output
func main() {
	configPath := flag.String("config", "/etc/ssh/sshd_config", "Path to the daemon configuration file") //Default: SSH Daemon Config File

	flag.Parse()

	fmt.Println("Fetching active services...\n")
	fmt.Println(getActiveServices())

	fmt.Println("\nFetching startup settings for services...\n")
	fmt.Println(getServiceStartupSettings())

	fmt.Println("\nFetching configuration...\n")
	fmt.Println(readDaemonConf(*configPath))
}
