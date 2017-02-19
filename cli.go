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
		Long: pharse + "\n" + logo,
		Run: func(cmd *cobra.Command, args []string) {
			c.prompt.listCommands()
		},
	}
}

func (c *cli) listEnvCMD() *cobra.Command {
	return &cobra.Command{
		Use:   "listenv [container name]",
		Short: "List all environment variables in the container",
		Long:  "List all environment variables in based on the input container name.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				printError(errors.New("listEnv only takes one container name"))
			}
			input := args[0]
			envs, err := c.commands.listEnv(input)
			printError(err)
			for k, v := range envs {
				fmt.Println(k + "=" + v)
			}
		},
	}
}

func (c *cli) listCMD() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all available containers",
		Long:  "List all available environment containers in the storage directory.",
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
		Long:  "Launch a process with all the environment variables sourced from the container.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				printError(errors.New("source only takes one container name"))
			}
			input := args[0]
			if !c.commands.checkFileExists(input) {
				printError(errors.New("container does not exist with the following name: " + input))
			}
			printError(c.commands.source(input))
		},
	}
}

func (c *cli) createCMD() *cobra.Command {
	createCMD := &cobra.Command{
		Use:     "create [container name] [environment variables, i.e name=value]",
		Short:   "create a container",
		Long:    "Create a container with the environment variables of your choice.",
		Example: "envcon create github test=test hello=world secret='skdfj8rfsnfsnfsdf'",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				printError(errors.New("create must take in a filename and at least one environment variable"))
			}
			name, envs := args[0], args[1:]
			if c.commands.checkFileExists(name) {
				printError(errors.New("container already exists with the following name: " + name))
			}
			envMap := make(map[string]string)
			for k := range envs {
				keyValues := strings.Split(envs[k], envSep)
				if len(keyValues) != 2 {
					printError(errors.New("the following environment was incorrent: " + envs[k]))
				}
				envMap[keyValues[0]] = keyValues[1]
			}
			printError(c.commands.create(name, envMap, cliEncrypt))
		},
	}
	createCMD.Flags().BoolVarP(&cliEncrypt, "encrypt", "e", false, "encrypt the file")
	return createCMD
}

func (c *cli) updateCMD() *cobra.Command {
	return &cobra.Command{
		Use:     "update [container name] [environment variables, i.e name=value]",
		Short:   "update a container",
		Long:    "update a container with the environment variables of your choice.",
		Example: "envcon update github test=test hello=world secret='skdfj8rfsnfsnfsdf'",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				printError(errors.New("update must take in a filename and at least one environment variable"))
			}
			name, envs := args[0], args[1:]
			if !c.commands.checkFileExists(name) {
				printError(errors.New("container does not exist with the following name: " + name))
			}
			envMap := make(map[string]string)
			for k := range envs {
				keyValues := strings.Split(envs[k], envSep)
				if len(keyValues) != 2 {
					printError(errors.New("the following environment was incorrent: " + envs[k]))
				}
				envMap[keyValues[0]] = keyValues[1]
			}
			printError(c.commands.update(name, envMap))
		},
	}
}

func (c *cli) deleteCMD() *cobra.Command {
	return &cobra.Command{
		Use:   "delete [container name]",
		Short: "delete a container",
		Long:  "delete a container based on the container name",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				printError(errors.New("delete only takes one container name"))
			}
			name := args[0]
			if !c.commands.checkFileExists(name) {
				printError(errors.New("container does not exist with the following name: " + name))
			}
			printError(c.commands.delete(name))
		},
	}
}
