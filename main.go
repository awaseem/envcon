package main

import (
	"os"
	"os/exec"
)

var langs = []string{
	"c",
	"c++",
	"lua",
	"go",
	"js",
	"ruby",
	"python",
}

func main() {
	// wd, err := os.Getwd()
	// must(err)
	// fs := fileStorage{
	// 	concel:        &aesCryp{},
	// 	storageFolder: wd + "/data",
	// }
	// envMap := make(map[string]string)
	// envMap["hello"] = "worldksldfskldfjsklfjsdklfjsdklfjlfksjfkldsjfsklfjdsklfjsfkljsklfjl"
	// envMap["hello2"] = "worldksldfskldfjsklfjsdklfjsdklfjlfksjfkldsjfsklfjdsklfjsfkljsklfjl"
	// envMap["hello3"] = "worldksldfskldfjsklfjsdklfjsdklfjlfksjfkldsjfsklfjdsklfjsfkljsklfjl"
	// envMap["hello4"] = "worldksldfskldfjsklfjsdklfjsdklfjlfksjfkldsjfsklfjdsklfjsfkljsklfjl"
	// envMap["hello5"] = "worldksldfskldfjsklfjsdklfjsdklfjlfksjfkldsjfsklfjdsklfjsfkljsklfjl"
	// f, err := fs.getFile("test.json")
	// must(err)
	// f.setContent(envMap, "")
	// f.save()
	// f.close()

	// wd, err := os.Getwd()
	// must(err)
	// fs := fileStorage{
	// 	concel:        &aesCryp{},
	// 	storageFolder: wd,
	// }
	// f, err := fs.getFile("test.json")
	// mapEnv, err := f.getContent("")
	// must(err)
	// fmt.Println(mapEnv["hello"])

	wd, err := os.Getwd()
	must(err)
	fs := &fileStorage{
		concel:        &aesCryp{},
		storageFolder: wd + "/data",
	}

	p := &interactivePrompt{
		fileStore: fs,
		session:   &session{},
	}
	p.listCommands()
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
