package merkle

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// define a struct of type Transaction that has a sender, receiver, and amount
type Transaction struct {
	Sender   string
	Receiver string
	Amount   int
}

// function to hash a transaction using the SHA256 algorithm
func (t *Transaction) Hash() string {
	// concatenate the transaction fields into a single string
	transactionString := t.Stringify()

	// hash the transaction string using the SHA256 algorithm
	hash := sha256.Sum256([]byte(transactionString))

	// return the hash as string
	return hex.EncodeToString(hash[:])
}

// function to stringify a transaction
func (t *Transaction) Stringify() string {
	return t.Sender + " -> " + t.Receiver + ": " + fmt.Sprintf("%d", t.Amount)
}

// function to create a new transaction
func NewTransaction(sender string, receiver string, amount int) *Transaction {
	return &Transaction{
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
	}
}
