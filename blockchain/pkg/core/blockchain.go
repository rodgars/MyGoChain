package core

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Blockchain represents the entire blockchain.
type Blockchain struct {
	Chain           []Block
	TransactionPool []Transaction
}

// SaveBlockchain saves the blockchain to a file.
func (bc *Blockchain) SaveBlockchain(filename string) error {
	data, err := json.MarshalIndent(bc, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

// LoadBlockchain loads the blockchain from a file.
func (blockchain *Blockchain) LoadBlockchain(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error loading blockchain from file:", err)
		return
	}

	var bcFile Blockchain
	err = json.Unmarshal(data, &bcFile)
	if err != nil {
		fmt.Println("Error converting json into blockchain:", err)
		return
	}

	blockchain.Chain = bcFile.Chain
	blockchain.TransactionPool = bcFile.TransactionPool
}

func (blockchain *Blockchain) InitiateBlockChain(fileName string) {

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		// Create a new blockchain with the genesis block
		genesisBlock := NewBlock(0, []Transaction{}, "")
		blockchain.Chain = []Block{*genesisBlock}

		// Save blockchain to file
		err = blockchain.SaveBlockchain(fileName)
		if err != nil {
			fmt.Println("Error saving blockchain:", err)
			return
		}
	} else {
		// Load existing blockchain from file
		blockchain.LoadBlockchain(fileName)
	}
}

// PrintBlockchain prints the current state of the blockchain.
func (bc *Blockchain) PrintBlockchain() {
	fmt.Println("Blockchain:")

	// Display blockchain
	for _, block := range bc.Chain {
		fmt.Printf("Index: %d\nTimestamp: %v\nPrevHash: %s\nHash: %s\nNonce: %d\n\n",
			block.Index, block.Timestamp, block.PrevHash, block.Hash, block.Nonce)

		fmt.Printf("\tTRANSACTIONS\n\t----------------------\n")
		for _, tran := range block.Data {
			fmt.Printf("\tSender: %s\n\tRecipient: %s\n\tAmount: %f\n\tTimestamp: %v\n\tSignature: %s\n", tran.Sender, tran.Recipient, tran.Amount, tran.Timestamp, base64.StdEncoding.EncodeToString(tran.Signature))
			fmt.Printf("\t----------------------\n\n")
		}
	}
}

// PrintTransactionPool prints the current state of the transaction pool.
func (bc *Blockchain) PrintTransactionPool() {
	fmt.Println("Transaction Pool:")
	for _, transaction := range bc.TransactionPool {
		fmt.Printf("Sender: %s\nRecipient: %s\nAmount: %f\nTimestamp: %s\n\n",
			transaction.Sender, transaction.Recipient, transaction.Amount, transaction.Timestamp)
	}
}

func (bc *Blockchain) CreateTransactionInPool(sender, recipient string, amount float64) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}

	transaction := CreateTransaction(sender, recipient, amount, privateKey.PublicKey)
	signature, err := SignTransaction(transaction, privateKey)
	if err != nil {
		fmt.Println("Error signing transaction:", err)
		return
	}
	transaction.Signature = signature

	bc.TransactionPool = append(bc.TransactionPool, transaction)

	bc.CreateBlockFromTransactionPool()
}

func (bc *Blockchain) CreateBlockFromTransactionPool() {
	if len(bc.TransactionPool) >= 3 {
		newBlock := Block{
			Index:     len(bc.Chain),
			Timestamp: time.Now(),
			Data:      bc.TransactionPool,
			PrevHash:  bc.Chain[len(bc.Chain)-1].Hash,
			Nonce:     0,
		}

		newBlock.MineBlock()

		bc.Chain = append(bc.Chain, newBlock)

		bc.TransactionPool = nil
	}
}
