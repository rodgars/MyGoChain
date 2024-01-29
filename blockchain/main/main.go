package main

import (
	"log"
	"os"

	"github.com/rodgars/MyGoChain/pkg/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Blockchain CLI"
	app.Usage = "Interact with the blockchain through a CLI"

	cliHandler := &cmd.CLIHandler{}
	cliHandler.InitiateBlockChain("data/blockchain.json")

	app.Commands = cliHandler.GetCommands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
