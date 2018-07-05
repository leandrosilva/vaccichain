package main

import (
	"fmt"
)

func main() {
	config := LoadConfig()
	fmt.Println(config)

	gblock := InitiateBlockchain()
	fmt.Println("Genesis Block:", gblock)
}
