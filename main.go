package main

import (
	"os"
)

func main() {
	// Check if a command is provided, immediately exit if not.
	if len(os.Args) < 2 {
		logger.Warn("Expected a command - type `zenforge help` to get options.")
		os.Exit(0)
	}

	// Dispatch the command
	switch os.Args[1] {
	case "init":
		command.Init()
	case "new":
		command.New()
	case "demo":
		command.Demo()
	case "build":
		command.Build()
	case "preview":
		command.Preview()
	case "help":
		command.Help()
	default:
		logger.Error("Unknown command: %s\n", os.Args[1])
	}

}
