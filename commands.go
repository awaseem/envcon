package main

import (
	"strings"
)

const (
	envconFileExt = ".json"
)

type commands struct {
	fileStore filer
	session   launcher
	input     inputer
}

func (p *commands) checkFileExists(fileName string) bool {
	return p.fileStore.exists(fileName + envconFileExt)
}

func (p *commands) list() ([]string, error) {
	fileNames := []string{}
	files, err := p.fileStore.listFiles()
	if err != nil {
		return nil, err
	}
	for i := range files {
		fileNames = append(fileNames, strings.Replace(files[i], envconFileExt, "", -1))
	}
	return fileNames, nil
}

func (p *commands) listEnv(fileName string) (map[string]string, error) {
	var pass string
	file, err := p.fileStore.getFile(fileName + envconFileExt)
	if err != nil {
		return nil, err
	}
	if file.fileContent.Encrypted {
		pass = p.input.passwordMasked("File is encrypted, please enter the passpharse")
	}
	return file.getContent(pass)
}

func (p *commands) source(fileName string) error {
	var pass string
	var envs map[string]string
	envFile, err := p.fileStore.getFile(fileName + envconFileExt)
	if err != nil {
		return err
	}
	if envFile.fileContent.Encrypted {
		pass = p.input.passwordMasked("File is encrypted, please enter the passpharse")
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
		pass = p.input.passwordMasked("Enter a passpharse")
	}
	envFile, err := p.fileStore.newFile(fileName+envconFileExt, encrypted)
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
	envFile, err := p.fileStore.getFile(fileName + envconFileExt)
	if err != nil {
		return err
	}
	if envFile.fileContent.Encrypted {
		pass = p.input.passwordMasked("File is encrypted, please enter the passpharse")
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
	return p.fileStore.deleteFile(fileName + envconFileExt)
}
