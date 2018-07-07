package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/leandrosilva/vaccichain/protocol"
	"google.golang.org/grpc"
)

func main() {
	config := LoadConfig()
	fmt.Println(config)

	blockchain := NewBlockchain()
	fmt.Println("Blockchain ID:", blockchain.ID)
	fmt.Println("Genesis Block:", blockchain.GetGenesisBlock())

	listener, err := net.Listen("tcp", "127.0.0.1:"+config.RPC.Port)
	if err != nil {
		log.Fatalf("Failed to listen on port %v: %v", config.RPC.Port, err)
	}

	server := grpc.NewServer()
	pb.RegisterBlockchainServiceServer(server, &RPCServer{})
	server.Serve(listener)
}
