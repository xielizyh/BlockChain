package main

import (
	"BitCoin/blockchain"
	"fmt"
	"time"
)

func main() {
	// 创建区块链
	bc := blockchain.NewBlockChain()
	// 添加区块，即交易
	bc.AddBlock("Send 1 BTC to David")
	bc.AddBlock("Send 2 more BTC to David")

	// 打印区块链信息
	for _, b := range bc.Blocks {
		fmt.Printf("Now| timestamp: %v\n", time.Unix(b.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Printf("Now| data: %s\n", b.Data)
		fmt.Printf("Pre| hash: %x\n", b.PreBlockHash)
		fmt.Printf("Now| hash: %x\n", b.Hash)
		fmt.Println()
	}
}
