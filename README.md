# Merkle Tree Builder

This is a simple implementation of a Merkle Tree builder for a given set of transactions. This builder uses `goroutines` to parallelize the process.
## Usage
To use the Merkle tree, you can create a slice of transactions and build the tree using the `BuildFromTransactions` function:
```go
import (
    "fmt"
    "github.com/naveed949/merkle-builder/merkle"
)

func main() {
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

    // build the Merkle tree from the transactions
    merkleRoot := merkle.BuildFromTransactions(transactions)

    // print the Merkle tree root
    fmt.Println("Merkle tree root:", merkleRoot)
}
```
You can also build the tree using goroutines and channels with the `BuildFromTransactionsConcurrent` function:
```go
// build the Merkle tree from the transactions using goroutines and channels
merkleRootConcurrent := merkle.BuildFromTransactionsConcurrent(transactions)

// print the Merkle tree root
fmt.Println("Merkle tree root (concurrent):", merkleRootConcurrent)
```

### Contributing
Contributions are welcome! Please open an issue or submit a pull request.

### License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
