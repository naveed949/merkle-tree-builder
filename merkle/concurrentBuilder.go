package merkle

import (
	"crypto/sha256"
	"encoding/hex"
)

// define struct of type Hash that has a value of hash and index of hash to ensure order is maintained upon receiving from channel
type Hash struct {
	Value string
	Index int
}

// function to construct a Merkle tree from a slice of hashes using goroutines and channels
func BuildConcurrent(hashes []string) []string {
	// create a channel of strings to hold the hashes
	hashChannel := make(chan Hash) // unbuffered channel

	// if the number of hashes is odd, duplicate the last hash
	if len(hashes)%2 != 0 {
		hashes = append(hashes, hashes[len(hashes)-1])
	}

	// iterate over the hashes in steps of 2
	for i := 0; i < len(hashes); i += 2 {
		index := i
		if i != 0 {
			index = i / 2
		}
		// launch a goroutine to concatenate the two hashes
		go func(firstHash string, secondHash string, index int) {
			// concatenate the two hashes
			concatenatedHashes := firstHash + secondHash

			// hash the concatenated hashes using the SHA256 algorithm
			hash := sha256.Sum256([]byte(concatenatedHashes))
			// send the hash of hashes to the channel
			hashChannel <- Hash{Value: hex.EncodeToString(hash[:]), Index: index}
		}(hashes[i], hashes[i+1], index)
	}

	// create a slice of strings to hold the new hashes
	var newHashes []string = make([]string, len(hashes)/2)

	// iterate over the hashes in steps of 2
	for i := 0; i < len(hashes); i += 2 {
		// receive the hash of hashes from the channel
		hash := <-hashChannel

		// append the hash to the slice of hashes
		newHashes[hash.Index] = hash.Value
	}

	// if the number of hashes is greater than 1, recursively call the function
	if len(newHashes) > 1 {
		return BuildConcurrent(newHashes) // recursive call to construct next level of tree
	}

	// return the slice of hashes as the Merkle tree root
	return newHashes
}

// function to construct a Merkle tree from a slice of transactions using goroutines and channels
func BuildFromTransactionsConcurrent(transactions []*Transaction) []string {
	// create a channel of strings to hold the hashes
	hashChannel := make(chan Hash) // unbuffered channel

	// iterate over the transactions
	for index, transaction := range transactions {
		// launch a goroutine to hash the transaction using the SHA256 algorithm
		go func(transaction *Transaction, index int) {
			hashChannel <- Hash{Value: transaction.Hash(), Index: index}
		}(transaction, index)
	}

	// create a slice of strings to hold the hashes
	var hashes []string = make([]string, len(transactions))

	// iterate over the transactions
	for range transactions {
		// receive the hash from the channel
		hash := <-hashChannel

		// append the hash to the slice of hashes
		hashes[hash.Index] = hash.Value
	}

	// construct the Merkle tree from the hashes
	return BuildConcurrent(hashes)
}
