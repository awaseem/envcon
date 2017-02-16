package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

type cli struct {
	commands commander
	prompt   prompter
}

func (c *cli) rootCMD() *cobra.Command {
	return &cobra.Command{
		Use:   "envcon",
		Short: "Envcon executable enviroment vairable containers",
		Long:  "Envcon allows you to create enviroment variable containers and launch a process within that container.",
		Run: func(cmd *cobra.Command, args []string) {
			c.prompt.listCommands()
		},
	}
}

func (c *cli) listCMD() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all available containers",
		Long:  "List all available enviroment conatiners in the storage directory",
		Run: func(cmd *cobra.Command, args []string) {
			files, err := c.commands.list()
			printError(err)
			for k := range files {
				fmt.Println(files[k])
			}
		},
	}
}
