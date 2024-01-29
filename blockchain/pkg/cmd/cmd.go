package cmd

import (
	"fmt"

	"github.com/rodgars/MyGoChain/pkg/core"
	"github.com/urfave/cli"
)

// CLIHandler handles CLI commands for interacting with the blockchain.
type CLIHandler struct {
	blockchain *core.Blockchain
	filename   string
}

func (ch *CLIHandler) InitiateBlockChain(filename string) {
	ch.filename = filename
	blockchain := &core.Blockchain{}
	blockchain.InitiateBlockChain(ch.filename)

	ch.blockchain = blockchain
}

// CreateTransaction creates a new transaction and adds it to the blockchain.
func (ch *CLIHandler) CreateTransaction(sender, recipient string, amount float64) {
	ch.blockchain.CreateTransactionInPool(sender, recipient, amount)
	// Save blockchain to file
	err := ch.blockchain.SaveBlockchain(ch.filename)
	if err != nil {
		fmt.Println("Error saving blockchain:", err)
		return
	}
}

// CreateTransaction creates a new transaction and adds it to the blockchain.
func (ch *CLIHandler) PrintBlockchain() {
	ch.blockchain.PrintBlockchain()
}

// CreateTransaction creates a new transaction and adds it to the blockchain.
func (ch *CLIHandler) PrintTransactionPool() {
	ch.blockchain.PrintTransactionPool()
}

func (ch *CLIHandler) GetCommands() []cli.Command {
	return []cli.Command{
		{
			Name:  "bc",
			Usage: "Check the current status of the blockchain",
			Action: func(c *cli.Context) {
				ch.PrintBlockchain()
			},
		},
		{
			Name:  "tp",
			Usage: "Check the current status of the transaction pool",
			Action: func(c *cli.Context) {
				ch.PrintTransactionPool()
			},
		},
		{
			Name:  "transaction",
			Usage: "Create a new transaction and add it to the blockchain",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "sender",
					Usage: "Sender of the transaction",
				},
				cli.StringFlag{
					Name:  "recipient",
					Usage: "Recipient of the transaction",
				},
				cli.Float64Flag{
					Name:  "amount",
					Usage: "Amount of the transaction",
				},
			},
			Action: func(c *cli.Context) {
				sender := c.String("sender")
				recipient := c.String("recipient")
				amount := c.Float64("amount")

				if sender == "" || recipient == "" || amount == 0 {
					fmt.Println("Invalid transaction parameters.")
					return
				}

				ch.CreateTransaction(sender, recipient, amount)
			},
		},
	}
}
