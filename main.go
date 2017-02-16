package main

const (
	storageFolderLoc = "/usr/local/var/envcon"
)

func main() {
	fs := &fileStorage{
		concel:        &aesCryp{},
		storageFolder: storageFolderLoc,
	}
	c := &commands{
		fileStore: fs,
		session:   &session{},
		input:     &userInput{},
	}
	p := &interactivePrompt{
		commands: c,
		input:    &userInput{},
	}
	cli := &cli{
		commands: c,
		prompt:   p,
	}
	// setup data storage if folder does not exist
	err := fs.createStore()
	printError(err)
	// setup prompts and command line
	var rootCmd = cli.rootCMD()
	rootCmd.AddCommand(cli.listCMD())
	rootCmd.AddCommand(cli.sourceCMD())
	rootCmd.AddCommand(cli.listEnvCMD())
	rootCmd.AddCommand(cli.createCMD())
	rootCmd.AddCommand(cli.updateCMD())
	rootCmd.AddCommand(cli.deleteCMD())
	rootCmd.Execute()
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
