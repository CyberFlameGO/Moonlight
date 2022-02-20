//go:build mage
// +build mage

package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

var Default = Build

// Build is TODO
func Build() error {
	fmt.Println("Building...")
	if err := checkDeps("go", "mvn", "java"); err != nil {
		return err
	}
	return nil
}

// Run is TODO
func Run() {
	fmt.Println("Running...")
}

// Clean is TODO
func Clean() {
	fmt.Println("Cleaning...")
}

// Helper functions

func cmd(name string, args ...string) error {
	return exec.Command(name, args...).Run()
}

func where(p string) error {
	if runtime.GOOS == "windows" {
		return cmd("where", p)
	} else {
		return cmd("command", "-v", p)
	}
}

func checkDeps(files ...string) error {
	for _, file := range files {
		fmt.Printf("checking if %s is installed...", file)
		if err := where(file); err != nil {
			fmt.Println("no")
			return err
		}
		fmt.Println("yes")
	}
	return nil
}
