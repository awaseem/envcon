package main

import "os"

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
	wd, err := os.Getwd()
	must(err)
	fs := &fileStorage{
		concel:        &aesCryp{},
		storageFolder: wd + "/data",
	}

	c := &commands{
		fileStore: fs,
		session:   &session{},
		input:     &userInput{},
	}
	p := &interactivePrompt{
		commands: c,
		input:    &userInput{},
	}
	p.listCommands()
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
