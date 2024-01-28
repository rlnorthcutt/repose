package main

import (
	"fmt"
)

// Defining a new public type 'Template'
type Filesystem int

// Defining a global varaiable for Template
var filesystem Filesystem

// **********  Template methods  ****************************************
func (f *Filesystem) Test() {
	fmt.Printf("Filesystem.go present and accounted for.")
}
