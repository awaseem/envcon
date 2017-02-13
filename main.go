package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	wd, err := os.Getwd()
	must(err)
	fs := fileStorage{
		settings: &settings{
			encrypt: true,
			key:     []byte("12345thisisatest"),
		},
		concel:        &aesCryp{},
		storageFolder: wd,
	}
	f, err := fs.getFile("test.envcon")
	must(err)
	fmt.Println(f.get("hello"))
	must(f.close())
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
