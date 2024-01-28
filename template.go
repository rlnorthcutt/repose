package main

import (
	"fmt"
)

// Defining a new public type 'Template'
type Template int

// Defining a global varaiable for Template
var template Template

// **********  Template methods  ****************************************
func (t *Template) Test() {
	fmt.Printf("Template.go present and accounted for.")
}
