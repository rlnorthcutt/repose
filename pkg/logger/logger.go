package logger

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Defining a new public type 'Logger'
type Logger struct {
	isVerbose bool
}

// ANSI color codes
const (
	// Reset to default
	Reset = "\033[0m"

	// Special
	Bold      = "\033[1m"
	Underline = "\033[4m"

	// Text colors
	Black  = "\033[30m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

func New(isVerbose bool) *Logger {
	return &Logger{
		isVerbose: isVerbose,
	}
}

// DETAIL method for the logger type (verbose)
// This method formats and prints a pain message without color
func (l *Logger) Detail(message string, value ...any) {
	if l.isVerbose {
		tag := "------- "
		fmt.Printf(tag+message+"\n", value...)
	}
}

// INFO method for the logger type
// This method formats and prints an info message with cyan color
func (l *Logger) Info(message string, value ...any) {
	tag := Cyan + "[INFO] " + Reset
	fmt.Printf(tag+message+"\n", value...)
}

// WARN method for the logger type
// This method formats and prints a warning message with yellow color
func (l *Logger) Warn(message string, value ...any) {
	tag := Yellow + "[WARNING] " + Reset
	fmt.Printf(tag+message+"\n", value...)
}

// ERR method for the logger type
// This method formats and prints an error message with red color
func (l *Logger) Error(message string, value ...any) {
	tag := Red + "[ERROR] " + Reset
	fmt.Printf(tag+message+"\n", value...)
}

// FATAL method for the logger type
// This method formats and prints a fatal message and exits the program
func (l *Logger) Fatal(message string, value ...any) {
	tag := Bold + Purple + "[FATAL] " + Reset
	fmt.Printf("\n"+tag+message+"\n", value...)
	os.Exit(1)
}

// SUCCESS method for the logger type
// This method formats and prints a success message with green color
func (l *Logger) Success(message string, value ...any) {
	tag := Bold + Green + "[SUCCESS] " + Reset
	fmt.Printf(tag+message+"\n\n", value...)
}

// DEBUG method for the logger type
// This method formats and prints a debug message with red color and exits the program
func (l *Logger) Debug(message string, value ...any) {
	tag := Bold + Red + "[--dEbUg--] " + Reset
	fmt.Printf(tag+message+"\n", value...)
	os.Exit(1)
}

// Helper function to prompt for input in a standard way
func (l *Logger) PromptForInput(prompt, defaultValue string) string {
	reader := bufio.NewReader(os.Stdin)

	// Use this instead of global.log.Info to avoid the newline character
	fmt.Printf("------- %s [%s]: ", prompt, defaultValue)

	input, err := reader.ReadString('\n')
	if err != nil {
		l.Error("Error reading input:", err)
		return defaultValue
	}

	input = strings.TrimSpace(input)
	if input == "" {
		return defaultValue
	}

	return input
}
