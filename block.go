package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"
)

var difficulty = 1

// GetDifficulty returns the current difficulty level for miners
func GetDifficulty() int {
	return difficulty
}

// ChangeDifficulty stabilishes a new level of difficulty for mining
func ChangeDifficulty(newDifficulty int) {
	difficulty = newDifficulty
}

// Block representation
type Block struct {
	Index      int
	Timestamp  string
	Data       string
	Hash       string
	PrevHash   string
	Difficulty int
	Nonce      string
}

// IsBlockValid checks if a block is valid or what
func (b *Block) IsBlockValid(previousBlock Block) bool {
	if b.Index != previousBlock.Index+1 {
		return false
	}

	if b.PrevHash != previousBlock.Hash {
		return false
	}

	if b.Hash != calculateHash(*b) {
		return false
	}

	return true
}

// NewGenesisBlock generates the very first block of the blockchain
func NewGenesisBlock() Block {
	var genesisBlock Block

	genesisBlock.Index = 0
	genesisBlock.Timestamp = time.Now().String()
	genesisBlock.Data = "Genesis Block"
	genesisBlock.PrevHash = ""
	genesisBlock.Difficulty = GetDifficulty()
	genesisBlock.Nonce = ""
	genesisBlock.Hash = calculateHash(genesisBlock)

	return genesisBlock
}

// NewBlock generates a new block and mines it
func NewBlock(previousBlock Block, data string) (Block, error) {
	var newBlock Block

	newBlock.Index = previousBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Data = data
	newBlock.PrevHash = previousBlock.Hash
	newBlock.Difficulty = GetDifficulty()

	if !isCandidateValid(newBlock, previousBlock) {
		return newBlock, errors.New("Candidate block is not valid")
	}

	mineBlock(&newBlock)

	return newBlock, nil
}

func mineBlock(newBlock *Block) {
	for i := 0; ; i++ {
		hex := fmt.Sprintf("%x", i)
		newBlock.Nonce = hex

		hash := calculateHash(*newBlock)
		if isHashValid(hash, newBlock.Difficulty) {
			newBlock.Hash = hash
			break
		}
	}
}

func isCandidateValid(candidateBlock Block, previousBlock Block) bool {
	if candidateBlock.Index != previousBlock.Index+1 {
		return false
	}

	if candidateBlock.PrevHash != previousBlock.Hash {
		return false
	}

	if candidateBlock.Data == "" {
		return false
	}

	return true
}

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + block.Data + block.PrevHash + block.Nonce
	hash := sha256.New()
	hash.Write([]byte(record))
	hashed := hash.Sum(nil)

	return hex.EncodeToString(hashed)
}

func isHashValid(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix)
}
