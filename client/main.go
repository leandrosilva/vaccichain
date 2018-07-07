package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/leandrosilva/vaccichain/protocol"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to %v", err)
	}
	defer conn.Close()

	client := pb.NewBlockchainServiceClient(conn)
	res, err := client.GetGenesisBlock(context.Background(), &pb.GenesisBlockRequest{})
	if err != nil {
		log.Fatalf("Failed calling RPCServer: %v", err)
	}

	fmt.Println("Genesis Block:", res.GenesisBlock)
}
