package merkle

import (
	"crypto/sha256"
	"encoding/hex"
)

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
