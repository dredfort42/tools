package logprinter

import (
	"fmt"
	"time"
)

// Color definitions
const (
	RED    = "\033[1;31m"
	GREEN  = "\033[1;32m"
	YELLOW = "\033[1;33m"
	RESET  = "\033[0m"
)

// Get the current time in the "2024-02-01 11:01:10" format
func getCurrentTime() string {
	return time.Now().Format("2024-02-01 11:01:10")
}

// PrintSuccess prints a success message
func PrintSuccess(msg string, info string) {
	if info == "" {
		fmt.Printf("%s %s[S] %s%s\n", getCurrentTime(), GREEN, msg, RESET)
	} else {
		fmt.Printf("%s %s[S] %s: %s%s\n", getCurrentTime(), GREEN, msg, info, RESET)
	}
}

// PrintWarning prints a warning message
func PrintWarning(msg string, info string) {
	if info == "" {
		fmt.Printf("%s %s[W] %s%s\n", getCurrentTime(), YELLOW, msg, RESET)
	} else {
		fmt.Printf("%s %s[W] %s: %s%s\n", getCurrentTime(), YELLOW, msg, info, RESET)
	}
}

// PrintError prints an error message
func PrintError(msg string, err error) {
	if err == nil {
		fmt.Printf("%s %s[E] %s%s\n", getCurrentTime(), RED, msg, RESET)
	} else {
		fmt.Printf("%s %s[E] %s: %s%s\n", getCurrentTime(), RED, msg, err.Error(), RESET)
	}
}

// PrintInfo prints an info message
func PrintInfo(msg string, info string) {
	if info == "" {
		fmt.Printf("%s [I] %s\n", getCurrentTime(), msg)
	} else {
		fmt.Printf("%s [I] %s: %s\n", getCurrentTime(), msg, info)
	}
}
