package main

import (
	"errors"
	"time"
)

// Blockchain represents the current state of the blockchain
type Blockchain struct {
	ID     string
	blocks []Block
}

// NewBlockchain creates a brand new blockchain with a genesis block
func NewBlockchain() Blockchain {
	gblock := NewGenesisBlock()
	id := CalculateHash(time.Now().String())
	blocks := []Block{gblock}
	blockchain := Blockchain{id, blocks}

	return blockchain
}

// GetBlocks returns a copy of the current state of the blockchain
func (bc *Blockchain) GetBlocks() []Block {
	return bc.blocks
}

// GetGenesisBlock returns the first block in the chain
func (bc *Blockchain) GetGenesisBlock() Block {
	return bc.blocks[0]
}

// GetLatestBlock returns the latest block in the chain
func (bc *Blockchain) GetLatestBlock() Block {
	return bc.blocks[len(bc.blocks)-1]
}

// GetBlockCount return the height of the current state
func (bc *Blockchain) GetBlockCount() int {
	return len(bc.blocks)
}

// AddBlock appends a new block to the blockchain
func (bc *Blockchain) AddBlock(newBlock Block) error {
	if !newBlock.IsBlockValid(bc.GetLatestBlock()) {
		return errors.New("Invalid block cannot be added to the blockchain")
	}
	bc.blocks = append(bc.blocks, newBlock)

	return nil
}

// GetBlockByIndex returns a given block by its index
func (bc *Blockchain) GetBlockByIndex(index int) (Block, error) {
	if int(index) < 1 || int(index) >= len(bc.blocks) {
		return Block{}, errors.New("Block index out of blockchain height")
	}

	return bc.blocks[index], nil
}

// GetBlockByHash returns a given block by its hash
func (bc *Blockchain) GetBlockByHash(hash string) (Block, error) {
	if len(hash) == 0 {
		return Block{}, errors.New("Empty hash? Are you kidding me?")
	}

	for i := 0; i < len(bc.blocks); i++ {
		if bc.blocks[i].Hash == hash {
			return bc.blocks[i], nil
		}
	}

	return Block{}, errors.New("Block not found")
}

// GetBlockRange returns a range of block starting from the given index provided that the genesis block is right
func (bc *Blockchain) GetBlockRange(genesisBlockHash string, startingIndex int) ([]Block, error) {
	if genesisBlockHash != bc.GetGenesisBlock().Hash {
		return []Block{}, errors.New("Invalid genesis block: " + genesisBlockHash)
	}

	if startingIndex < 0 || startingIndex >= bc.GetBlockCount() {
		return []Block{}, errors.New("Invalid starting index: " + string(startingIndex))
	}

	var blocks []Block
	for i := startingIndex; i < bc.GetBlockCount(); i++ {
		blocks = append(blocks, bc.blocks[i])
	}

	return blocks, nil
}

// ReplaceChain overrides the current state by the new one
func (bc *Blockchain) ReplaceChain(givenChain []Block) {
	if isValidChain(*bc, givenChain) {
		bc.blocks = givenChain
	}
}

func isValidChain(blockchain Blockchain, givenChain []Block) bool {
	if givenChain[0].Data != blockchain.GetGenesisBlock().Data {
		return false
	}

	if len(givenChain) <= blockchain.GetBlockCount() {
		return false
	}

	for i := 1; i < len(givenChain); i++ {
		if !givenChain[i].IsBlockValid(givenChain[i-1]) {
			return false
		}
	}

	return true
}
