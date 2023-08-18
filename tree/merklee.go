package tree

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

// function to construct a Merkle tree from a slice of hashes
func Build(hashes []string) []string {
	// create a slice of strings to hold the hashes
	var newHashes []string

	// if the number of hashes is odd, duplicate the last hash
	if len(hashes)%2 != 0 {
		hashes = append(hashes, hashes[len(hashes)-1])
	}

	// iterate over the hashes in steps of 2
	for i := 0; i < len(hashes); i += 2 {
		// concatenate the two hashes
		concatenatedHashes := hashes[i] + hashes[i+1]

		// hash the concatenated hashes using the SHA256 algorithm
		hash := sha256.Sum256([]byte(concatenatedHashes))

		// append the hash to the slice of hashes
		newHashes = append(newHashes, hex.EncodeToString(hash[:]))
	}

	// if the number of hashes is greater than 1, recursively call the function
	if len(newHashes) > 1 {
		return Build(newHashes) // recursive call to construct next level of tree
	}

	// return the slice of hashes as the Merkle tree root
	return newHashes
}

// function to construct a Merkle tree from a slice of transactions
func BuildFromTransactions(transactions []*Transaction) []string {
	// create a slice of strings to hold the hashes
	var hashes []string

	// iterate over the transactions
	for _, transaction := range transactions {
		// hash the transaction using the SHA256 algorithm
		hashes = append(hashes, transaction.Hash())
	}

	// construct the Merkle tree from the hashes
	return Build(hashes)
}
