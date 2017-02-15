package main

import "fmt"

var (
	promptCommands = []string{"source", "create", "update", "delete"}
)

type interactivePrompt struct {
	commands commander
}

func (p *interactivePrompt) listCommands() {
	printLogo()
	i := inputChoose("Please select from the following commands", promptCommands)
	switch promptCommands[i] {
	case "source":
		p.source()
	case "create":
		p.create()
	case "update":
		p.update()
	case "delete":
		p.delete()
	}
}

func (p *interactivePrompt) source() {
	files, err := p.commands.list()
	printError(err)
	i := inputChoose("Select from the following enviroments", files)
	envFile := files[i]
	p.commands.source(envFile)
}

func (p *interactivePrompt) create() {
	var done bool
	envs := make(map[string]string)
	fileName := inputStringRequired("Enter a name for this container")
	encrypted := inputConfirm("Would you like to encrypt this container?(Yes,y/No,n)")
	for !done {
		key := inputStringRequired("Enter a key")
		value := inputStringRequired("Enter a value")
		envs[key] = value
		done = inputConfirm("stop adding enviroment variables?(Yes,y/No,n)")
	}
	printError(p.commands.create(fileName, envs, encrypted))
}

func (p *interactivePrompt) update() {
	var done bool
	envs := make(map[string]string)
	files, err := p.commands.list()
	printError(err)
	i := inputChoose("Pick a container to update", files)
	envFile := files[i]
	for !done {
		key := inputStringRequired("Enter a key")
		fmt.Println("printing the key" + key)
		value := inputStringRequired("Enter a value")
		envs[key] = value
		done = inputConfirm("stop adding/updating variables enviroment variables?(Yes,y/No,n)")
	}
	printError(p.commands.update(envFile, envs))
}

func (p *interactivePrompt) delete() {
	files, err := p.commands.list()
	printError(err)
	i := inputChoose("Pick a container to delete", files)
	printError(p.commands.delete(files[i]))
}
