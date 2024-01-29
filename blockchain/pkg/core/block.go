package core

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

// Block represents a block in the blockchain.
type Block struct {
	Index     int
	Timestamp time.Time
	Data      []Transaction
	PrevHash  string
	Hash      string
	Nonce     int // Used for proof-of-work
}

// CalculateHash calculates the hash of the block.
func (b *Block) CalculateHash() {
	// Concatenate block fields and calculate SHA-256 hash
	hashDataInput := ""
	for _, t := range b.Data {
		hashDataInput += fmt.Sprintf("%s%s%f%d%s", t.Sender, t.Recipient, t.Amount, t.Timestamp.Unix(), t.Signature)
	}
	hashInput := fmt.Sprintf("%d%d%s%s%d", b.Index, b.Timestamp.Unix(), hashDataInput, b.PrevHash, b.Nonce)
	hash := sha256.New()
	hash.Write([]byte(hashInput))
	b.Hash = hex.EncodeToString(hash.Sum(nil))
}

// MineBlock finds a valid hash by adjusting the nonce.
func (b *Block) MineBlock() {
	// Define the target difficulty (number of leading zeros in the hash)
	targetDifficulty := 4
	zeroString := strings.Repeat("0", targetDifficulty)

	for {
		b.CalculateHash()

		// Check if the calculated hash meets the target difficulty
		if b.Hash[:targetDifficulty] == zeroString {
			break
		}

		// Increment the nonce and try again
		b.Nonce++
	}
}

// NewBlock creates a new block with the given data and previous hash.
func NewBlock(index int, data []Transaction, prevHash string) *Block {
	block := &Block{
		Index:     index,
		Timestamp: time.Now(),
		Data:      data,
		PrevHash:  prevHash,
		Nonce:     0,
	}

	// Mine the block by finding a valid hash (proof-of-work)
	block.MineBlock()

	return block
}
