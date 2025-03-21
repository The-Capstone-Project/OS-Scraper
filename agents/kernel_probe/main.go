package kernel_probe

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

// Function to execute a shell command and return its output
func executeCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

// Function to list loaded kernel modules
func listLoadedModules() (string, error) {
	return executeCommand("lsmod")
}

// Function to get details about available kernel modules
func getModuleInfo(moduleName string) (string, error) {
	return executeCommand("modinfo", moduleName)
}

// Function to read configuration files for kernel modules
func readConfigFiles() (map[string]string, error) {
	configFiles := []string{"/etc/modprobe.d/", "/etc/modules-load.d/"}
	configData := make(map[string]string)

	for _, dir := range configFiles {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			return nil, err
		}

		for _, file := range files {
			if !file.IsDir() {
				content, err := ioutil.ReadFile(dir + file.Name())
				if err != nil {
					return nil, err
				}
				configData[dir+file.Name()] = string(content)
			}
		}
	}
	return configData, nil
}

func Run() {
	// List loaded kernel modules
	loadedModules, err := listLoadedModules()
	if err != nil {
		fmt.Println("Error listing loaded modules:", err)
		return
	}
	fmt.Println("Loaded Kernel Modules:")
	fmt.Println(loadedModules)

	// Get details about a specific kernel module (example: "ext4")
	moduleInfo, err := getModuleInfo("ext4")
	if err != nil {
		fmt.Println("Error getting module info:", err)
		return
	}
	fmt.Println("Module Info for 'ext4':")
	fmt.Println(moduleInfo)

	// Read configuration files for kernel modules
	configFiles, err := readConfigFiles()
	if err != nil {
		fmt.Println("Error reading config files:", err)
		return
	}
	fmt.Println("Kernel Module Configuration Files:")
	for file, content := range configFiles {
		fmt.Printf("File: %s\nContent:\n%s\n", file, content)
	}
}
