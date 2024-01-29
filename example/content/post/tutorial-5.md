---
title: "Go Tutorial - Building a Static Site Generator - Part 5"
description: "Exploring the beauty and simplicity of static websites in the modern web era."
tags: [static sites, web development, SiteStat]
image: /images/zen-static.jpg
index: true
publish: true
author: "Ron Northcutt"
publish_date: "2024-01-30"
template: "blog-post.tmpl"
---

# Go Tutorial - Building a Static Site Generator - Part  5
## Build out logger
We will use this for outputting responses

```
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
// This method formats and prints an info message with green color
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
	os.Exit(0)
}

```