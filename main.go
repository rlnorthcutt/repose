package main

import (
	"flag"
	"os"

	"github.com/rlnorthcutt/repose/internal/config"
	"github.com/rlnorthcutt/repose/pkg/logger"
)

// Func main should be as small as possible and do as little as possible by convention
func main() {
	// Parse the command line arguments
	rootPath, isVerbose, args := parse()

	// Create the logger instance to use
	logger := logger.New(isVerbose)

	// Check the environment - returns if the config file is not found
	checkEnvironment(args, logger, rootPath)

	// Dispatch the command
	command.Dispatch(args, logger, rootPath)
}

// **********  Private Main methods  **********

// Parse the command for flags and arguments
func parse() (rootPath string, isVerbose bool, args []string) {
	// Define flags for verbose mode and root path with both long and shorthand versions.
	flag.BoolVar(&isVerbose, "verbose", false, "Display detailed logger messages")
	flag.BoolVar(&isVerbose, "v", false, "Display detailed logger messages (shorthand)")
	flag.StringVar(&rootPath, "root", ".", "Root path of the project")
	flag.StringVar(&rootPath, "r", ".", "Root path of the project (shorthand)")

	// Parse the CLI flags
	flag.Parse()

	return rootPath, isVerbose, flag.Args()
}

// Check the environment for the required files and directories & load the config file
func checkEnvironment(args []string, logger *logger.Logger, rootPath string) {
	// Check if a command is provided, immediately exit if not.
	if len(args) < 1 {
		logger.Warn("Expected a command - use \033[1;96m`repose help`\033[0m to get options.")
		os.Exit(0)
	}

	command := args[0]

	// Check if the config file exists
	if command != "init" && command != "help" {
		logger.Detail("Checking if config file exists.")

		// Setup a new config struct
		configurator := config.New(logger)

		if !configurator.Exists(rootPath) {
			logger.Warn("No config file found. You need to run \033[1;96m`repose init`\033[0m first.")
			os.Exit(0)
		}
	}
}
