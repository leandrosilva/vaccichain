package main

import (
	"errors"
)

var blockchain []Block

// InitiateBlockchain creates the genesis block, adds it, and returns it after all
func InitiateBlockchain() Block {
	gblock := NewGenesisBlock()
	blockchain = append(blockchain, gblock)

	return gblock
}

// GetBlockchain returns a copy of the current state of the blockchain
func GetBlockchain() []Block {
	return blockchain
}

// GetGenesisBlock returns the first block in the chain
func GetGenesisBlock() Block {
	return blockchain[0]
}

// GetLatestBlock returns the latest block in the chain
func GetLatestBlock() Block {
	return blockchain[len(blockchain)-1]
}

// GetBlockCount return the height of the current state
func GetBlockCount() int {
	return len(blockchain)
}

// AddBlock appends a new block to the blockchain
func AddBlock(newBlock Block) error {
	if !newBlock.IsBlockValid(GetLatestBlock()) {
		return errors.New("Invalid block cannot be added to the blockchain")
	}
	blockchain = append(blockchain, newBlock)

	return nil
}

// ReplaceChain overrides the current state by the new one
func ReplaceChain(givenBlockchain []Block) {
	if isValidChain(givenBlockchain) {
		blockchain = givenBlockchain
	}
}

func isValidChain(givenBlockchain []Block) bool {
	if givenBlockchain[0].Data != GetGenesisBlock().Data {
		return false
	}

	if len(givenBlockchain) <= GetBlockCount() {
		return false
	}

	for i := 1; i < len(givenBlockchain); i++ {
		if !givenBlockchain[i].IsBlockValid(givenBlockchain[i-1]) {
			return false
		}
	}

	return true
}
