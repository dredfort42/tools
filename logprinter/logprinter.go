package logprinter

import (
	"fmt"
	"os"
	"time"
)

// Color definitions
const (
	RED    = "\033[1;31m"
	GREEN  = "\033[1;32m"
	BLUE   = "\033[1;34m"
	YELLOW = "\033[1;33m"
	RESET  = "\033[0m"
)

// Get the current time in the "2006-01-02 15:04:05" format
func getCurrentTime() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05")
}

// Success prints a success message
func Success(msg string, info ...string) {
	fmt.Printf("%s %s[S] %s", getCurrentTime(), GREEN, msg)

	for i := 0; i < len(info); i++ {
		fmt.Printf(" | %s", info[i])
	}

	fmt.Printf("%s\n", RESET)
}

// Warning prints a warning message
func Warning(msg string, info ...string) {
	fmt.Printf("%s %s[W] %s", getCurrentTime(), YELLOW, msg)

	for i := 0; i < len(info); i++ {
		fmt.Printf(" | %s", info[i])
	}

	fmt.Printf("%s\n", RESET)
}

// Error prints an error message
func Error(msg string, err error) {
	fmt.Printf("%s %s[E] %s", getCurrentTime(), RED, msg)

	if err != nil {
		fmt.Printf(" | %s", err.Error())
	}

	fmt.Printf("%s\n", RESET)
}

// Info prints an info message
func Info(msg string, info ...string) {
	fmt.Printf("%s [I] %s", getCurrentTime(), msg)

	for i := 0; i < len(info); i++ {
		fmt.Printf(" | %s", info[i])
	}

	fmt.Println()
}

// Debug prints a debug message if the DEBUG environment variable is set to "1"
func Debug(msg string, info ...string) {
	if os.Getenv("DEBUG") != "1" {
		return
	}

	fmt.Printf("%s %s[D] %s%s", getCurrentTime(), BLUE, RESET, msg)

	for i := 0; i < len(info); i++ {
		fmt.Printf(" | %s", info[i])
	}

	fmt.Println()
}
