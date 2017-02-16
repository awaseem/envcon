package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

const (
	envSep = "="
)

var (
	cliEncrypt = false
)

type cli struct {
	commands commander
	prompt   prompter
}

func (c *cli) rootCMD() *cobra.Command {
	return &cobra.Command{
		Use:  "envcon",
		Long: logo,
		Run: func(cmd *cobra.Command, args []string) {
			c.prompt.listCommands()
		},
	}
}

func (c *cli) listCMD() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all available containers",
		Long:  "List all available enviroment conatiners in the storage directory.",
		Run: func(cmd *cobra.Command, args []string) {
			files, err := c.commands.list()
			printError(err)
			for k := range files {
				fmt.Println(files[k])
			}
		},
	}
}

func (c *cli) sourceCMD() *cobra.Command {
	return &cobra.Command{
		Use:   "source [container to run]",
		Short: "source a container",
		Long:  "Launch a process with all the enviroment variables sourced from the container.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				printError(errors.New("source only takes one container name"))
			}
			input := args[0]
			printError(c.commands.source(input))
		},
	}
}

func (c *cli) createCMD() *cobra.Command {
	return &cobra.Command{
		Use:   "create [container name] [enviroment variables, i.e name=value]",
		Short: "create a container",
		Long:  "Create a container with the enviroment variables of your choice.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				printError(errors.New("create must take in a filename and at least one enviroment variable"))
			}
			name, envs := args[0], args[1:]
			envMap := make(map[string]string)
			for k := range envs {
				keyValues := strings.Split(envs[k], envSep)
				if len(keyValues) != 2 {
					printError(errors.New("the following enviroment was incorrent"))
				}
				envMap[keyValues[0]] = keyValues[1]
			}
			printError(c.commands.create(name, envMap, cliEncrypt))
		},
	}
}
