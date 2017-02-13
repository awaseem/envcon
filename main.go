package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	//create new file and save
	// wd, err := os.Getwd()
	// must(err)
	// fs := fileStorage{
	// 	concel:        &aesCryp{},
	// 	storageFolder: wd,
	// }
	// envMap := make(map[string]string)
	// envMap["hello"] = "world"
	// f, err := fs.newFile("test.json", false)
	// must(err)
	// f.setContent(envMap, "test")
	// f.save()

	wd, err := os.Getwd()
	must(err)
	fs := fileStorage{
		concel:        &aesCryp{},
		storageFolder: wd,
	}
	f, err := fs.getFile("test.json")
	mapEnv, err := f.getContent("")
	must(err)
	fmt.Println(mapEnv["hello"])
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
