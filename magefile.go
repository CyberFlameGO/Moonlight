//go:build mage
// +build mage

package main

import (
	"fmt"
	"os/exec"
)

var Default = Build

// Build is TODO
func Build() error {
	fmt.Println("Building...")
	cmd := exec.Command("go", "build", "-o", "MyApp", ".")
	return cmd.Run()
}

// Run is TODO
func Run() {
	fmt.Println("Running...")
}

// Clean is TODO
func Clean() {
	fmt.Println("Cleaning...")
}
