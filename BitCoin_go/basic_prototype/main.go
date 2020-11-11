package main

import "fmt"

func main() {
	// 创建区块链
	bc := NewBlockChain()

	bc.AddBlock("Send 1 BTC to xieli")
	bc.AddBlock("Send 2 more BTC to xieli")
	for _, block := range bc.blocks {
		fmt.Printf("Prev.hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
