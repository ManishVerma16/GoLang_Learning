package main

import (
	"fmt"
	"go_blockchain/blockchain"
)



func main()  {
	fmt.Println("Hello World")

	chain := blockchain.InitBlockChain()
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Manish Verma")
	chain.AddBlock("Lalit Munna Gupta")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

	}
}
