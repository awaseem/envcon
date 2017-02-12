package main

import (
	"os"
	"os/exec"
)

func main() {
	envs := make(map[string]string)
	envs["CLIENT_ID"] = "OH MY GOD"
	s := &session{}
	s.launch(envs)
}

func child() {
	cmd := exec.Command("sh")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
