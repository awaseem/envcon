package main

import prompt "github.com/segmentio/go-prompt"

var (
	promptCommands = []string{"source", "create", "update", "delete"}
)

type interactivePrompt struct {
	commands commander
}

func (p *interactivePrompt) listCommands() {
	printLogo()
	i := prompt.Choose("Please select from the following commands", promptCommands)
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
	i := prompt.Choose("Select from the following enviroments", files)
	envFile := files[i]
	p.commands.source(envFile)
}

func (p *interactivePrompt) create() {
	var done bool
	envs := make(map[string]string)
	fileName := prompt.StringRequired("Enter a name for this container")
	encrypted := prompt.Confirm("Would you like to encrypt this container?(Yes,y/No,n)")
	for !done {
		key := prompt.StringRequired("Enter a key")
		value := prompt.StringRequired("Enter a value")
		envs[key] = value
		done = prompt.Confirm("stop adding enviroment variables?(Yes,y/No,n)")
	}
	printError(p.commands.create(fileName, envs, encrypted))
}

func (p *interactivePrompt) update() {
	var done bool
	envs := make(map[string]string)
	files, err := p.commands.list()
	printError(err)
	i := prompt.Choose("Pick a container to update", files)
	envFile := files[i]
	for !done {
		key := prompt.StringRequired("Enter a key")
		value := prompt.StringRequired("Enter a value")
		envs[key] = value
		done = prompt.Confirm("stop adding/updating variables enviroment variables?(Yes,y/No,n)")
	}
	printError(p.commands.update(envFile, envs))
}

func (p *interactivePrompt) delete() {
	files, err := p.commands.list()
	printError(err)
	i := prompt.Choose("Pick a container to delete", files)
	printError(p.commands.delete(files[i]))
}
