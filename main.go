package main

import (
	"fmt"
)

func main() {
	config := LoadConfig()
	fmt.Println(config)

	gblock := CreateBlockchain()
	fmt.Println("Genesis Block:", gblock)
}
