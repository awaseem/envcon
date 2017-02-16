package main

import "fmt"

var (
	promptCommands = []string{"list", "listenv", "source", "create", "update", "delete"}
)

type interactivePrompt struct {
	commands commander
	input    inputer
}

func (p *interactivePrompt) listCommands() {
	printLogo()
	i := p.input.choose("Please select from the following commands", promptCommands)
	switch promptCommands[i] {
	case "list":
		p.list()
	case "listenv":
		p.listEnv()
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

func (p *interactivePrompt) list() {
	files, err := p.commands.list()
	printError(err)
	for i := range files {
		fmt.Println(files[i])
	}
}

func (p *interactivePrompt) listEnv() {
	files, err := p.commands.list()
	printError(err)
	i := p.input.choose("Select from the following enviroments", files)
	envFile := files[i]
	envMap, err := p.commands.listEnv(envFile)
	printError(err)
	for k, v := range envMap {
		fmt.Println(k + "=" + v)
	}
}

func (p *interactivePrompt) source() {
	files, err := p.commands.list()
	printError(err)
	i := p.input.choose("Select from the following enviroments", files)
	envFile := files[i]
	p.commands.source(envFile)
}

func (p *interactivePrompt) create() {
	var done bool
	envs := make(map[string]string)
	fileName := p.input.stringRequired("Enter a name for this container")
	encrypted := p.input.confirm("Would you like to encrypt this container?(Yes,y/No,n)")
	for !done {
		key := p.input.stringRequired("Enter a key")
		value := p.input.stringRequired("Enter a value")
		envs[key] = value
		done = p.input.confirm("stop adding enviroment variables?(Yes,y/No,n)")
	}
	printError(p.commands.create(fileName, envs, encrypted))
}

func (p *interactivePrompt) update() {
	var done bool
	envs := make(map[string]string)
	files, err := p.commands.list()
	printError(err)
	i := p.input.choose("Pick a container to update", files)
	envFile := files[i]
	for !done {
		key := p.input.stringRequired("Enter a key")
		fmt.Println("printing the key" + key)
		value := p.input.stringRequired("Enter a value")
		envs[key] = value
		done = p.input.confirm("stop adding/updating variables enviroment variables?(Yes,y/No,n)")
	}
	printError(p.commands.update(envFile, envs))
}

func (p *interactivePrompt) delete() {
	files, err := p.commands.list()
	printError(err)
	i := p.input.choose("Pick a container to delete", files)
	printError(p.commands.delete(files[i]))
}
