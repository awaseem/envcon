package main

import prompt "github.com/segmentio/go-prompt"

type commands struct {
	fileStore filer
	session   launcher
}

func (p *commands) list() ([]string, error) {
	return p.fileStore.listFiles()
}

func (p *commands) source(fileName string) error {
	var pass string
	var envs map[string]string
	envFile, err := p.fileStore.getFile(fileName)
	if err != nil {
		return err
	}
	if envFile.fileContent.Encrypted {
		pass = prompt.PasswordMasked("File is encrypted, please enter the passpharse")
	}
	envs, err = envFile.getContent(pass)
	if err != nil {
		return err
	}
	p.session.launch(envs)
	return nil
}

func (p *commands) create(fileName string, env map[string]string, encrypted bool) error {
	var pass string
	if encrypted {
		pass = prompt.PasswordMasked("Enter a passpharse")
	}
	envFile, err := p.fileStore.newFile(fileName+".json", encrypted)
	if err != nil {
		return err
	}
	err = envFile.setContent(env, pass)
	if err != nil {
		return err
	}
	return envFile.save()
}

func (p *commands) update(fileName string, env map[string]string) error {
	var pass string
	envs := make(map[string]string)
	envFile, err := p.fileStore.getFile(fileName)
	if err != nil {
		return err
	}
	if envFile.fileContent.Encrypted {
		pass = prompt.PasswordMasked("File is encrypted, please enter the passpharse")
	}
	envs, err = envFile.getContent(pass)
	if err != nil {
		return err
	}
	for k, v := range env {
		envs[k] = v
	}
	err = envFile.setContent(envs, pass)
	if err != nil {
		return err
	}
	return envFile.save()
}

func (p *commands) delete(fileName string) error {
	return p.fileStore.deleteFile(fileName)
}
