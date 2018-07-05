package main

import (
	"fmt"
)

func main() {
	config := LoadConfig()
	fmt.Println(config)

	blockchain := NewBlockchain()
	fmt.Println("Blockchain ID:", blockchain.ID)
	fmt.Println("Genesis Block:", blockchain.GetGenesisBlock())
}
