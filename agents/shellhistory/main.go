package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	flag.IntVar(&n, "n", -1, "Number of history items to fetch")
	flag.Parse()

	zshHistoryFile := os.Getenv("HOME") + "/.zsh_history"

	if n == -1 {
		fmt.Print("Warning: Entire history will be collected! Are you sure? (yes/no): ")
	} else if n == 0 {
		fmt.Print("Warning: Entire history will be collected (n=0)! Are you sure? (yes/no): ")
	} else {
		fmt.Printf("Warning: Last %d history items will be collected! Are you sure? (yes/no): ", n)
	}

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	input = strings.TrimSpace(strings.ToLower(input))
	if input != "yes" && input != "y" {
		fmt.Println("Operation cancelled.")
		return
	}

	if n <= 0 {
		printEntireHistory(zshHistoryFile)
	} else {
		printLimitedHistory(zshHistoryFile, n)
	}
}

func printEntireHistory(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading .zsh_history file:", err)
		return
	}
	fmt.Println("Content of .zsh_history file:")
	fmt.Println(string(content))
}

func printLimitedHistory(filename string, n int) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading .zsh_history file:", err)
		return
	}
	lines := strings.Split(string(content), "\n")
	start := len(lines) - n
	if start < 0 {
		start = 0
	}
	fmt.Printf("Last %d lines of .zsh_history file:\n", n)
	for _, line := range lines[start:] {
		fmt.Println(line)
	}
}
