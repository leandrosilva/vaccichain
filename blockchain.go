package main

import (
	"errors"
)

var blockchain []Block

// CreateBlockchain creates the genesis block, adds it, and returns it after all
func CreateBlockchain() Block {
	gblock := GenerateGenesisBlock()
	blockchain = append(blockchain, gblock)

	return gblock
}

// GetBlockchain returns a copy of the current state of the blockchain
func GetBlockchain() []Block {
	return blockchain
}

// GetBlockCount return the length of the current state
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

// GetLatestBlock returns the latest block in the chain
func GetLatestBlock() Block {
	return blockchain[len(blockchain)-1]
}

// ReplaceChain overrides the current state by the new one
func ReplaceChain(newBlocks []Block) {
	if len(newBlocks) > len(blockchain) {
		blockchain = newBlocks
	}
}
