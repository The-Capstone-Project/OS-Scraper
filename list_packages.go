package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// RunCommand runs a shell command and returns the output or an error.
func RunCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s: %s", err, string(output))
	}
	return string(output), nil
}

// SuggestSolution provides installation instructions if a command is not found.
func SuggestSolution(command string) string {
	switch command {
	case "pip":
		return "pip not found. You can install Python and pip by visiting: https://www.python.org/downloads/"
	case "npm":
		return "npm not found. Install Node.js (which includes npm) from: https://nodejs.org/"
	case "gem":
		return "RubyGems not found. Install Ruby (which includes gem) from: https://www.ruby-lang.org/en/downloads/"
	case "composer":
		return "Composer not found. Install PHP and Composer from: https://getcomposer.org/download/"
	case "go":
		return "Go not found. Install Go from: https://golang.org/dl/"
	case "cargo":
		return "Cargo not found. Install Rust and Cargo from: https://www.rust-lang.org/tools/install"
	case "yarn":
		return "Yarn not found. Install Yarn from: https://yarnpkg.com/getting-started/install"
	default:
		return "Command not found. Please install the required tool."
	}
}

// ListPythonPackages lists globally installed Python packages.
func ListPythonPackages() {
	fmt.Println("Globally Installed Python Packages:")
	output, err := RunCommand("pip", "list")
	if err != nil {
		fmt.Println(SuggestSolution("pip"))
		return
	}
	fmt.Println(output)
}

// ListNodePackages lists globally installed Node.js packages.
func ListNodePackages() {
	fmt.Println("Globally Installed Node.js Packages:")
	output, err := RunCommand("npm", "list", "-g", "--depth=0")
	if err != nil {
		fmt.Println(SuggestSolution("npm"))
		return
	}
	fmt.Println(output)
}

// ListLocalNodePackages lists locally installed Node.js packages in a given project directory.
func ListLocalNodePackages(projectDir string) {
	fmt.Println("Locally Installed Node.js Packages in:", projectDir)
	cmd := exec.Command("npm", "list", "--depth=0")
	cmd.Dir = projectDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(SuggestSolution("npm"))
		return
	}
	fmt.Println(string(output))
}

// ListRubyGems lists globally installed Ruby gems.
func ListRubyGems() {
	fmt.Println("Globally Installed Ruby Gems:")
	output, err := RunCommand("gem", "list")
	if err != nil {
		fmt.Println(SuggestSolution("gem"))
		return
	}
	fmt.Println(output)
}

// ListComposerPackages lists globally installed PHP Composer packages.
func ListComposerPackages() {
	fmt.Println("Globally Installed PHP Composer Packages:")
	output, err := RunCommand("composer", "global", "show")
	if err != nil {
		fmt.Println(SuggestSolution("composer"))
		return
	}
	fmt.Println(output)
}

// ListGoPackages lists installed Go modules for the current project.
func ListGoPackages() {
	fmt.Println("Installed Go Modules:")
	output, err := RunCommand("go", "list", "-m", "all")
	if err != nil {
		fmt.Println(SuggestSolution("go"))
		return
	}
	fmt.Println(output)
}

// ListRustCrates lists globally installed Rust crates.
func ListRustCrates() {
	fmt.Println("Globally Installed Rust Crates:")
	output, err := RunCommand("cargo", "install", "--list")
	if err != nil {
		fmt.Println(SuggestSolution("cargo"))
		return
	}
	fmt.Println(output)
}

// ListYarnPackages lists globally installed Yarn packages.
func ListYarnPackages() {
	fmt.Println("Globally Installed Yarn Packages:")
	output, err := RunCommand("yarn", "global", "list")
	if err != nil {
		fmt.Println(SuggestSolution("yarn"))
		return
	}
	fmt.Println(output)
}

// ListLocalYarnPackages lists locally installed Yarn packages in a given project directory.
func ListLocalYarnPackages(projectDir string) {
	fmt.Println("Locally Installed Yarn Packages in:", projectDir)
	cmd := exec.Command("yarn", "list")
	cmd.Dir = projectDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(SuggestSolution("yarn"))
		return
	}
	fmt.Println(string(output))
}

func main() {
	// List globally installed Python packages.
	ListPythonPackages()

	// List globally installed Node.js packages.
	ListNodePackages()

	// Example: List locally installed Node.js packages in the current directory.
	projectDir, err := filepath.Abs(".")
	if err != nil {
		log.Println("Error getting project directory:", err)
		return
	}
	ListLocalNodePackages(projectDir)

	// List globally installed Ruby gems.
	ListRubyGems()

	// List globally installed PHP Composer packages.
	ListComposerPackages()

	// List installed Go modules.
	ListGoPackages()

	// List globally installed Rust crates.
	ListRustCrates()

	// List globally installed Yarn packages.
	ListYarnPackages()

	// Example: List locally installed Yarn packages in the current directory.
	ListLocalYarnPackages(projectDir)

	// Add more functions as needed to list packages from other languages/environments.
}
