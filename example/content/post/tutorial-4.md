---
title: "Go Tutorial - Building a Static Site Generator - Part 4"
description: "Exploring the beauty and simplicity of static websites in the modern web era."
tags: [static sites, web development, Repose]
image: /images/zen-static.jpg
index: true
publish: true
author: "Ron Northcutt"
publish_date: "2024-01-30"
template: "blog-post.tmpl"
---

# Go Tutorial - Building a Static Site Generator - Part  4
## Shell script for running multifile Go project

🧙🏾‍♂️: Writing a shell script to run all `.go` files in the current directory is a straightforward task. The script will use the `go run` command, which can compile and run Go programs. Typically, you would end the file in `.sh` to denote that it is a shell script. However, here we will skip that to make it simpler for testing purposes.

Also, we want to exclude test files AND make sure that we can call this from other directories and have it _only_ aggregate the Go files in the directory it is in.

Here's a simple script to achieve this:

### Shell Script: `gotest`

```bash
#!/bin/bash

# Get the directory where the script is located
script_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

# Find all .go files in the script's directory, excluding *_test.go files
go_files=$(find "$script_dir" -maxdepth 1 -name "*.go" ! -name "*_test.go")

# Run all .go files with any provided arguments
echo "-------------------------- "
echo "|--  Testing Repose  --| "
echo "-------------------------- "
go run $go_files $@ 

```

### How to Use the Script

1. **Make the Script Executable**: 
   First, you need to make this script executable. You can do this by running the following command in your terminal:

   ```bash
   chmod +x gotest
   ```

2. **Run the Script**: 
   Execute the script by typing `./gotest` in your terminal.

### Notes

- This script assumes that all `.go` files in the current directory are part of the same package and can be run together. If this is not the case, the script might not work as expected.
- The script uses `find` to locate `.go` files. It's set to only look in the current directory (`-maxdepth 1`).
- The script checks if there are no `.go` files and exits with a message if none are found.
$* in the echo statement prints all script arguments as a single string.
"$@" is used to pass all the script arguments to the go run command. It preserves the argument boundaries (e.g., new post title is passed as three separate arguments).

Remember, shell scripts can vary based on the shell environment and operating system. This script is written for a standard Unix-like shell (like Bash), commonly found in Linux and macOS environments.

If you have any specific requirements or face issues with this script, feel free to ask for further assistance!