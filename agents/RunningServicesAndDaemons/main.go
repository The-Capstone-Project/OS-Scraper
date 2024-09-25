package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func runCommand(cmd string, args ...string) (string, error) {
	out, err := exec.Command(cmd, args...).CombinedOutput()
	return string(out), err
}

func getActiveServices() string {
	output, err := runCommand("systemctl", "list-units", "--type=service", "--state=running")
	if err != nil {
		return fmt.Sprintf("Error fetching active services: %s\n", err)
	}
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
	fmt.Printf("Total Number Of Active Services: %d\n", (len(serviceNames) - 5))
	return strings.Join(serviceNames[1:len(serviceNames)-4], "\n")
}

func getServiceStartupSettings() string {
	output, err := runCommand("systemctl", "list-unit-files", "--type=service")
	if err != nil {
		return fmt.Sprintf("Error fetching service startup settings: %s\n", err)
	}
	var serviceStatuses []string
	scanner := bufio.NewScanner(strings.NewReader(output))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				serviceStatuses = append(serviceStatuses, fmt.Sprintf("%s: %s", parts[0], parts[1]))
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Sprintf("Error scanning output: %s\n", err)
	}
	fmt.Printf("Total Number Of Services: %d\n", (len(serviceStatuses) - 2))
	return strings.Join(serviceStatuses[1:len(serviceStatuses)-1], "\n")
}

func getService(serviceName string) string {
	output, err := runCommand("systemctl", "status", serviceName)
	if err != nil {
		exitError, ok := err.(*exec.ExitError)
		if ok {
			switch exitError.ExitCode() {
			case 3:
				return fmt.Sprintf("The service '%s' does not exist or is not loaded.\n", serviceName)
			case 4:
				return fmt.Sprintf("The service '%s' is not running.\n", serviceName)
			default:
				return fmt.Sprintf("Error querying service '%s': %s\nOutput: %s\n", serviceName, err, output)
			}
		}
		return fmt.Sprintf("Error querying service '%s': %s\nOutput: %s\n", serviceName, err, output)
	}
	return output
}

func getDaemonConfig(serviceName string) string {
	configLocations := []string{
		fmt.Sprintf("/etc/systemd/system/%s.service", serviceName),
		fmt.Sprintf("/lib/systemd/system/%s.service", serviceName),
		fmt.Sprintf("/usr/lib/systemd/system/%s.service", serviceName),
		fmt.Sprintf("/etc/%s/%s.conf", serviceName, serviceName),
		fmt.Sprintf("/etc/%s.conf", serviceName),
	}

	for _, location := range configLocations {
		content, err := readConfigFile(location)
		if err == nil {
			return fmt.Sprintf("Configuration for %s (from %s):\n\n%s", serviceName, location, content)
		}
	}

	return fmt.Sprintf("Could not find configuration file for %s", serviceName)
}

func readConfigFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func main() {
	serviceName := flag.String("service", "", "Name of the specific service to fetch")
	showConfig := flag.Bool("show-config", false, "Show the configuration for the service")
	flag.Parse()

	if *serviceName != "" {
		if *showConfig {
			fmt.Println(getDaemonConfig(*serviceName))
		} else {
			fmt.Printf("Fetching information for service: %s\n\n", *serviceName)
			fmt.Println(getService(*serviceName))
		}
	} else {
		fmt.Println("Fetching active services...\n")
		fmt.Println(getActiveServices())
		fmt.Println("\nFetching startup settings for services...\n")
		fmt.Println(getServiceStartupSettings())
	}
}
