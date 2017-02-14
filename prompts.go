package main

import prompt "github.com/segmentio/go-prompt"
import "fmt"

var (
	commands = []string{"source", "create", "update", "delete"}
)

type interactivePrompt struct {
	fileStore filer
	session   launcher
}

func (p *interactivePrompt) listCommands() {
	printLogo()
	i := prompt.Choose("Please select from the following commands", commands)
	switch commands[i] {
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
	var pass string
	var envs map[string]string
	files, err := p.fileStore.listFiles()
	printError(err)
	i := prompt.Choose("Select from the following enviroments", files)
	envFile, err := p.fileStore.getFile(files[i])
	printError(err)
	if envFile.fileContent.Encrypted {
		pass = prompt.PasswordMasked("File is encrypted, please enter the passpharse")
	}
	envs, err = envFile.getContent(pass)
	printError(err)
	p.session.launch(envs)
}

func (p *interactivePrompt) create() {
	var pass string
	var done bool
	envs := make(map[string]string)
	fileName := prompt.StringRequired("Enter a name for this container")
	encrypted := prompt.Confirm("Would you like to encrypt this container?(Yes,y/No,n)")
	if encrypted {
		pass = prompt.PasswordMasked("Enter a passpharse")
	}
	envFile, err := p.fileStore.newFile(fileName+".json", encrypted)
	printError(err)
	for !done {
		key := prompt.StringRequired("Enter a key")
		value := prompt.StringRequired("Enter a value")
		envs[key] = value
		done = prompt.Confirm("stop adding enviroment variables?(Yes,y/No,n)")
	}
	printError(envFile.setContent(envs, pass))
	printError(envFile.save())
}

func (p *interactivePrompt) update() {
	var pass string
	var done bool
	envs := make(map[string]string)
	files, err := p.fileStore.listFiles()
	printError(err)
	i := prompt.Choose("Pick a container to update", files)
	envFile, err := p.fileStore.getFile(files[i])
	printError(err)
	if envFile.fileContent.Encrypted {
		pass = prompt.PasswordMasked("File is encrypted, please enter the passpharse")
	}
	envs, err = envFile.getContent(pass)
	printError(err)
	for !done {
		for k, v := range envs {
			fmt.Println(k + "=" + v)
		}
		key := prompt.StringRequired("Enter a key")
		value := prompt.StringRequired("Enter a value")
		envs[key] = value
		done = prompt.Confirm("stop adding/updating variables enviroment variables?(Yes,y/No,n)")
	}
	printError(envFile.setContent(envs, pass))
	printError(envFile.save())
}

func (p *interactivePrompt) delete() {
	files, err := p.fileStore.listFiles()
	printError(err)
	i := prompt.Choose("Pick a container to delete", files)
	printError(p.fileStore.deleteFile(files[i]))
}
