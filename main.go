package main

import (
	"os"
	"os/exec"
)

func main() {
	path, err := os.Getwd()
	must(err)
	f := &fileStorage{
		storageFolder: path,
	}
	must(err)
	// err = eFile.set("test1", "tester1")
	// err = eFile.set("test2", "tester2")
	// err = eFile.set("test3", "tester3")
	// err = eFile.set("test4", "tester4")
	// must(err)
	// must(eFile.save())
	// must(eFile.close())
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
