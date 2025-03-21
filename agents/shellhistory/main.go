package shellhistory

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func Run() {
	// Set default limit to 10
	n := flag.Int("n", 10, "Number of history items to fetch (default: 10)")
	flag.Parse()

	zshHistoryFile := os.Getenv("HOME") + "/.zsh_history"

	// If n is <= 0, default to 10
	if *n <= 0 {
		*n = 10
	}

	printLimitedHistory(zshHistoryFile, *n)
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
