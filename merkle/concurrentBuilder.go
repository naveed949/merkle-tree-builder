package merkle

import (
	"crypto/sha256"
	"encoding/hex"
)

// function to construct a Merkle tree from a slice of hashes using goroutines and channels
func BuildConcurrent(hashes []string) []string {
	// create a channel of strings to hold the hashes
	hashChannel := make(chan string) // unbuffered channel

	// if the number of hashes is odd, duplicate the last hash
	if len(hashes)%2 != 0 {
		hashes = append(hashes, hashes[len(hashes)-1])
	}

	// iterate over the hashes in steps of 2
	for i := 0; i < len(hashes); i += 2 {
		// launch a goroutine to concatenate the two hashes
		go func(firstHash string, secondHash string) { //TODO: which routine is executed first is not guaranteed
			// concatenate the two hashes
			concatenatedHashes := firstHash + secondHash

			// hash the concatenated hashes using the SHA256 algorithm
			hash := sha256.Sum256([]byte(concatenatedHashes))
			// send the hash of hashes to the channel
			hashChannel <- hex.EncodeToString(hash[:])
		}(hashes[i], hashes[i+1])
	}

	// create a slice of strings to hold the new hashes
	var newHashes []string

	// iterate over the hashes in steps of 2
	for i := 0; i < len(hashes); i += 2 {
		// receive the hash of hashes from the channel
		hashOfHashes := <-hashChannel

		// append the hash to the slice of hashes
		newHashes = append(newHashes, hashOfHashes)
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
	hashChannel := make(chan string) // unbuffered channel

	// iterate over the transactions
	for _, transaction := range transactions {
		// launch a goroutine to hash the transaction using the SHA256 algorithm
		go func(transaction *Transaction) {
			hashChannel <- transaction.Hash()
		}(transaction)
	}

	// create a slice of strings to hold the hashes
	var hashes []string

	// iterate over the transactions
	for range transactions {
		// receive the hash from the channel
		hash := <-hashChannel

		// append the hash to the slice of hashes
		hashes = append(hashes, hash)
	}

	// construct the Merkle tree from the hashes
	return BuildConcurrent(hashes)
}
