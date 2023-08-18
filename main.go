package main

import (
	"fmt"

	"github.com/naveed949/merklee-builder/tree"
)

func main() {
	fmt.Println("Hello, world!")
	// create a slice of transactions
	transactions := []*tree.Transaction{
		tree.NewTransaction("A", "B", 100),
		tree.NewTransaction("C", "D", 200),
		tree.NewTransaction("E", "F", 300),
		tree.NewTransaction("G", "H", 400),
		tree.NewTransaction("I", "J", 500),
		tree.NewTransaction("K", "L", 600),
		tree.NewTransaction("M", "N", 700),
		tree.NewTransaction("O", "P", 800),
	}

	// create a slice of hashes
	var hashes []string = make([]string, len(transactions))

	// iterate over the transactions and hash each one
	for i, transaction := range transactions {
		hashes[i] = transaction.Hash()
	}

	// build the Merkle tree
	merkleRoot := tree.Build(hashes)

	// print the Merkle tree root
	fmt.Println("Merkle tree root:", merkleRoot)

	// build the Merkle tree from the transactions
	merkleRootFromTransactions := tree.BuildFromTransactions(transactions)

	// print the Merkle tree root
	fmt.Println("Merkle tree root from transactions:", merkleRootFromTransactions)

}
