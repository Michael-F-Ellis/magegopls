//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/sh"
)

var Default = BuildAndRun

// BuildAndRun runs the magefile in mypkg/ to generate the package code. If that
// succeeds it builds and runs the application.
func BuildAndRun() {
	pwd, _ := os.Getwd()
	os.Chdir("mypkg")
	err := sh.Run("mage", "-v", "generate")
	if err != nil {
		fmt.Println("Generation failed!")
		return
	}
	os.Chdir(pwd)
	err = sh.Run("go", "build")
	if err != nil {
		fmt.Println("Build failed!")
		return
	}
	err = sh.Run("./magegopls")
	if err != nil {
		fmt.Println("Error running magegopls")
		return
	}
}
