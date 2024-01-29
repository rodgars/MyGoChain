# MyGoChain
Simple Blockchain implementation built entirely in GoLang.

### How to test

- install [GO](https://go.dev/doc/install) in your environment
- download this repository source code 
- open a terminal, navigate to the folder blockchain: `cd blockchain`
- execute the application: `go run .\main\main.go`

This will show all possible options to interact with the blockchain via this CLI:

```
NAME:
   Blockchain CLI - Interact with the blockchain through a CLI

USAGE:
   main.exe [global options] command [command options] [arguments...]

COMMANDS:
   bc           Check the current status of the blockchain
   tp           Check the current status of the transaction pool
   transaction  Create a new transaction and add it to the blockchain
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

- To print the blockchain, just `go run .\main\main.go bc`
- To print the transaction pool, just `go run .\main\main.go tp`
- To create a transaction (it first creates in the transaction pool), just `go run .\main\main.go transaction --sender=YOUR_SENDER_NAME --recipient=YOUR_RECIPIENT_NAME --amount=YOUR_AMOUNT`
  - One new block is mined after 3 transactions are added in the transaction pool
