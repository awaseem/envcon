package main

import (
	"os"
	"os/exec"
)

// terminal session with all arguments
type session struct{}

// launches a new session with envs, panics if error
func (s *session) launch(envs map[string]string) {
	for k, v := range envs {
		err := os.Setenv(k, v)
		if err != nil {
			panic(err)
		}
	}
	// set indicator to let user know that a session with custom envs has started
	err := os.Setenv("PS1", os.Getenv("PS1")+"(evncon) ")
	if err != nil {
		panic(err)
	}
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "sh"
	}
	cmd := exec.Command(shell)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
