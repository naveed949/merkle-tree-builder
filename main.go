package main

import (
	"fmt"

	"github.com/naveed949/merklee-builder/merkle"
)

func main() {
	fmt.Println("Hello, world!")
	// create a slice of transactions
	transactions := []*merkle.Transaction{
		merkle.NewTransaction("A", "B", 100),
		merkle.NewTransaction("C", "D", 200),
		merkle.NewTransaction("E", "F", 300),
		merkle.NewTransaction("G", "H", 400),
		merkle.NewTransaction("I", "J", 500),
		merkle.NewTransaction("K", "L", 600),
		merkle.NewTransaction("M", "N", 700),
		merkle.NewTransaction("O", "P", 800),
	}

	// create a slice of hashes
	var hashes []string = make([]string, len(transactions))

	// iterate over the transactions and hash each one
	for i, transaction := range transactions {
		hashes[i] = transaction.Hash()
	}

	// build the Merkle tree
	merkleRoot := merkle.Build(hashes)

	// print the Merkle tree root
	fmt.Println("Merkle tree root:", merkleRoot)

	// build the Merkle tree from the transactions
	merkleRootFromTransactions := merkle.BuildFromTransactions(transactions)

	// print the Merkle tree root
	fmt.Println("Merkle tree root from transactions:", merkleRootFromTransactions)

}
