package main

import (
	"fmt"
	"os"
)

const (
	logo = `
   _____  ___   ___________  _  __
  / __/ |/ / | / / ___/ __ \/ |/ /
 / _//    /| |/ / /__/ /_/ /    / 
/___/_/|_/ |___/\___/\____/_/|_/

`
	pharse = "Development environment container"
)

// printLogo print the logo of this application
func printLogo() {
	fmt.Println(logo)
	fmt.Println(pharse)
}

// printError print the error and exit
func printError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
