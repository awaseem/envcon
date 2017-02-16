package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

type cli struct {
	commands commander
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
