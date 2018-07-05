package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockchainEmptyBeforeCreation(t *testing.T) {
	assert.Equal(t, GetBlockCount(), 0, "Blockchain should be empty")
}

func TestGenesisBlockCreation(t *testing.T) {
	gblock := InitiateBlockchain()

	assert.Equal(t, "Genesis Block", gblock.Data, "Failed to create genesis block")
	assert.Equal(t, 1, GetBlockCount(), "Blockchain should contain only the genesis block")
}

func TestAddNewBlock(t *testing.T) {
	gblock := InitiateBlockchain()

	nblock, err := NewBlock(gblock, "Block 1")
	assert.Nil(t, err)

	err = AddBlock(nblock)
	assert.Nil(t, err)

	assert.Equal(t, 2, GetBlockCount(), "Blockchain should contain 2 blocks")
}

func TestNotAddAnyBlockTwice(t *testing.T) {
	gblock := InitiateBlockchain()
	err := AddBlock(gblock)
	assert.NotNil(t, err)
}

func TestNotAcceptInvalidCandidateBlockWithEmptyData(t *testing.T) {
	gblock := InitiateBlockchain()
	_, err := NewBlock(gblock, "")

	assert.NotNil(t, err)
	assert.Equal(t, 1, GetBlockCount(), "Blockchain should have only the genesis block")
}

func TestNotAcceptInvalidCandidateBlockWithWrongPrevious(t *testing.T) {
	gblock := InitiateBlockchain()

	block1, err := NewBlock(gblock, "Block 1")
	assert.Nil(t, err)
	err = AddBlock(block1)
	assert.Nil(t, err)
	assert.Equal(t, 2, GetBlockCount(), "Blockchain should contain 2 blocks")

	_, err = NewBlock(gblock, "Block 2")
	assert.Nil(t, err)
	assert.Equal(t, 2, GetBlockCount(), "Blockchain should still contain 2 blocks")
}
