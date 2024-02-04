package main

import (
	"fmt"
	"os"
)

// Defining a new public type 'Logger'
type Logger int

// Defining a global varaiable for Logger
var logger Logger

// **********  Logger methods  ****************************************

// INFO method for the logger type
// This method formats and prints an info message with cyan color
func (l *Logger) Info(message string, value ...any) {
	tag := "\u001B[0;36m[INFO]\u001B[0;39m "
	fmt.Printf(tag+message+"\n", value...)
}

// WARN method for the logger type
// This method formats and prints a warning message with yellow color
func (l *Logger) Warn(message string, value ...any) {
	tag := "\u001B[0;33m[WARNING]\u001B[0;39m "
	fmt.Printf(tag+message+"\n", value...)
}

// ERR method for the logger type
// This method formats and prints an error message with red color
func (l *Logger) Error(message string, value ...any) {
	tag := "\u001B[0;35m[ERROR]\u001B[0;39m "
	fmt.Printf(tag+message+"\n", value...)
}

// FATAL method for the logger type
// This method formats and prints a fatal message and exits the program
func (l *Logger) Fatal(message string, value ...any) {
	tag := "\u001B[0;31m[FATAL]\u001B[0;39m "
	fmt.Printf(tag+message+"\n", value...)
	os.Exit(1)
}

// SUCCESS method for the logger type
// This method formats and prints a success message with green color
func (l *Logger) Success(message string, value ...any) {
	tag := "\u001B[1;32m[SUCCESS]\u001B[0;39m "
	fmt.Printf(tag+message+"\n", value...)
}
