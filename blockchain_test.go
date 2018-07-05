package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockchainCreation(t *testing.T) {
	bc := NewBlockchain()

	assert.Equal(t, "Genesis Block", bc.GetGenesisBlock().Data, "Failed to create genesis block")
	assert.Equal(t, 1, bc.GetBlockCount(), "Blockchain should contain only the genesis block")
}

func TestAddNewBlock(t *testing.T) {
	bc := NewBlockchain()
	pblock := bc.GetGenesisBlock()

	nblock, err := NewBlock(pblock, "Block 1")
	assert.Nil(t, err)

	err = bc.AddBlock(nblock)
	assert.Nil(t, err)

	assert.Equal(t, 2, bc.GetBlockCount(), "Blockchain should contain 2 blocks")
}

func TestNotAddAnyBlockTwice(t *testing.T) {
	bc := NewBlockchain()
	gblock := bc.GetGenesisBlock()
	err := bc.AddBlock(gblock)
	assert.NotNil(t, err)
}

func TestNotAcceptInvalidCandidateBlockWithEmptyData(t *testing.T) {
	bc := NewBlockchain()
	gblock := bc.GetGenesisBlock()
	_, err := NewBlock(gblock, "")

	assert.NotNil(t, err)
	assert.Equal(t, 1, bc.GetBlockCount(), "Blockchain should have only the genesis block")
}

func TestNotAcceptInvalidCandidateBlockWithWrongPrevious(t *testing.T) {
	bc := NewBlockchain()
	gblock := bc.GetGenesisBlock()

	block1, err := NewBlock(gblock, "Block 1")
	assert.Nil(t, err)
	err = bc.AddBlock(block1)
	assert.Nil(t, err)
	assert.Equal(t, 2, bc.GetBlockCount(), "Blockchain should contain 2 blocks")

	_, err = NewBlock(gblock, "Block 2")
	assert.Nil(t, err)
	assert.Equal(t, 2, bc.GetBlockCount(), "Blockchain should still contain 2 blocks")
}
