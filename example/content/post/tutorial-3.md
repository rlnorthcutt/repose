---
title: "Go Tutorial - Building a Static Site Generator - Part 3"
description: "Exploring the beauty and simplicity of static websites in the modern web era."
tags: [static sites, web development, Repose]
image: /images/zen-static.jpg
index: true
publish: true
author: "Ron Northcutt"
publish_date: "2024-01-30"
template: "blog-post.tmpl"
---

# Go Tutorial - Building a Static Site Generator - Part  3
Managing the project

- created as a module, so all files will be built as a single compilation target
- this allows us to divide up our command into multiple files for better organization and make it easier to build
- however, this requires us either do the build OR run all of the go files together in order to test.
- so we will create a shell script to let us do a live test as we do development

## Setup
- Initialize a new Golang project.
- Set up a version control system (e.g., Git).
- Create the example folder structure for testing (config, content, template, web).
- create the go file placeholders
    - each object has a demo method that prints out which class it is
    - main.go calls the demo method on each object
    - this will let us have a "working" framework that we can then build out
- create the shell script to run the go files 
- test the dev script
- test the build

## #Sample code for objects
A note about methods. If the method is lowercase, then it is "private", so they
won't be accessible inside other files. So, you need to use Capitalized methods
for anything that needs to be accessed in other files.

```go
package main

import (
	"fmt"
)

// Defining a new public type 'Template'
type Template int

// Defining a global varaiable for Template
var template Template

// **********  Template methods  **********
func (t *Template) Test() {
	fmt.Printf("Template.go present and accounted for.")
}

```
## Sample code for main controller
```
package main

import (
    fmt
)

func main() {
    fmt.Printf("Calling all systems - role call:")
	// Test that things are wired up properly
    command.Test()
	filesystem.Test()
    logger.Test()
	template.Test()
}

```