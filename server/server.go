package main

import (
	"context"
	"errors"

	pb "github.com/leandrosilva/vaccichain/protocol"
)

var _ = pb.BlockMessage{}

// RPCServer implements the vaccichain's Protobuf protocol
type RPCServer struct {
	blockchain Blockchain
}

// SetBlockchain defines the blockchain it will operate on
func (s *RPCServer) SetBlockchain(bc Blockchain) {
	s.blockchain = bc
}

// GetBlockchain returns the current blockchain it operates on
func (s *RPCServer) GetBlockchain() Blockchain {
	return s.blockchain
}

// GetBlocks return all blocks
func (s *RPCServer) GetBlocks(ctx context.Context, req *pb.BlocksRequest) (*pb.BlocksResponse, error) {
	blocks := s.blockchain.GetBlocks()

	res := &pb.BlocksResponse{Blocks: []*pb.BlockMessage{}}
	for i := 0; i < len(blocks); i++ {
		res.Blocks = append(res.Blocks, packBlockIntoMessage(blocks[i]))
	}

	return res, nil
}

// GetGenesisBlock returns the genesis block
func (s *RPCServer) GetGenesisBlock(ctx context.Context, req *pb.GenesisBlockRequest) (*pb.GenesisBlockResponse, error) {
	block := s.blockchain.GetGenesisBlock()
	res := &pb.GenesisBlockResponse{GenesisBlock: packBlockIntoMessage(block)}

	return res, nil
}

// GetLatestBlock returns the genesis block
func (s *RPCServer) GetLatestBlock(ctx context.Context, req *pb.LatestBlockRequest) (*pb.LatestBlockResponse, error) {
	block := s.blockchain.GetLatestBlock()
	res := &pb.LatestBlockResponse{LatestBlock: packBlockIntoMessage(block)}

	return res, nil
}

// GetBlockCount returns the height of the blockchain
func (s *RPCServer) GetBlockCount(ctx context.Context, req *pb.BlockCountRequest) (*pb.BlockCountResponse, error) {
	count := s.blockchain.GetBlockCount()
	res := &pb.BlockCountResponse{BlockCount: int32(count)}

	return res, nil
}

// AddBlock proposes a new block to be part of the blockchain
func (s *RPCServer) AddBlock(ctx context.Context, req *pb.NewBlockRequest) (*pb.NewBlockResponse, error) {
	newBlock := unpackBlockFromMessage(req.NewBlock)
	res := &pb.NewBlockResponse{Accepted: true, RejectReason: ""}

	err := s.blockchain.AddBlock(newBlock)
	if err != nil {
		res.Accepted = false
		res.RejectReason = err.Error()
	}

	return res, err
}

// GetBlock returns a given block by its index or hash
func (s *RPCServer) GetBlock(ctx context.Context, req *pb.BlockRequest) (*pb.BlockResponse, error) {
	var block Block
	var err error

	if req.BlockIndex > 0 {
		block, err = s.blockchain.GetBlockByIndex(int(req.BlockIndex))
	} else if req.BlockHash != "" {
		block, err = s.blockchain.GetBlockByHash(req.BlockHash)
	} else {
		err = errors.New("Should provide an index or hash to try to find a block")
	}

	if err != nil {
		return &pb.BlockResponse{}, err
	}

	res := &pb.BlockResponse{Block: packBlockIntoMessage(block)}

	return res, nil
}

// GetBlockRange returns a number of blocks from a given index
func (s *RPCServer) GetBlockRange(ctx context.Context, req *pb.BlockRangeRequest) (*pb.BlockRangeResponse, error) {
	blocks, err := s.blockchain.GetBlockRange(req.GenesisBlockHash, int(req.StartingIndex))
	if err != nil {
		return &pb.BlockRangeResponse{}, err
	}

	res := &pb.BlockRangeResponse{Blocks: []*pb.BlockMessage{}}
	for i := 0; i < len(blocks); i++ {
		res.Blocks = append(res.Blocks, packBlockIntoMessage(blocks[i]))
	}

	return res, nil
}

func packBlockIntoMessage(b Block) *pb.BlockMessage {
	return &pb.BlockMessage{
		Index:      int32(b.Index),
		Timestamp:  b.Timestamp,
		Data:       b.Data,
		Hash:       b.Hash,
		PrevHash:   b.PrevHash,
		Difficulty: int32(b.Difficulty),
		Nonce:      b.Nonce,
	}
}

func unpackBlockFromMessage(bm *pb.BlockMessage) Block {
	return Block{
		Index:      int(bm.Index),
		Timestamp:  bm.Timestamp,
		Data:       bm.Data,
		Hash:       bm.Hash,
		PrevHash:   bm.PrevHash,
		Difficulty: int(bm.Difficulty),
		Nonce:      bm.Nonce,
	}
}
