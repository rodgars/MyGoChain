package core

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"time"
)

// Transaction represents a transaction in the blockchain.
type Transaction struct {
	Sender    string
	Recipient string
	Amount    float64
	Timestamp time.Time
	Signature []byte
	PublicKey rsa.PublicKey
}

// transactionsToString converts transactions to a string for hashing.
func transactionsToString(transactions []Transaction) string {
	var result string
	for _, tx := range transactions {
		result += fmt.Sprintf("%s%s%f%s", tx.Sender, tx.Recipient, tx.Amount, tx.Timestamp.String())
	}
	return result
}

// CreateTransaction creates a new transaction.
func CreateTransaction(sender, recipient string, amount float64, publicKey rsa.PublicKey) Transaction {
	return Transaction{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
		Timestamp: time.Now(),
		PublicKey: publicKey,
	}
}

// SignTransaction signs a transaction with a private key.
func SignTransaction(tx Transaction, privateKey *rsa.PrivateKey) ([]byte, error) {
	hash := sha256.New()
	hash.Write([]byte(transactionsToString([]Transaction{tx})))

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash.Sum(nil))
	if err != nil {
		return nil, err
	}

	return signature, nil
}

// VerifyTransaction verifies the authenticity of a transaction's signature.
func VerifyTransaction(tx Transaction) bool {
	hash := sha256.New()
	hash.Write([]byte(transactionsToString([]Transaction{tx})))

	err := rsa.VerifyPKCS1v15(&tx.PublicKey, crypto.SHA256, hash.Sum(nil), tx.Signature)
	return err == nil
}
