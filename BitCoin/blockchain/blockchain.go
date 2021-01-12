package blockchain

import "BitCoin/block"

// BlockChain 区块链是一系列区块组成的链式结构
type BlockChain struct {
	Blocks []*block.Block
}

// NewBlockChain 创建一个区块链
func NewBlockChain() *BlockChain {
	// 最初的区块链只有一个创世区块
	return &BlockChain{[]*block.Block{block.NewGenesisBlock()}}
}

// AddBlock 向区块链中加入一个区块
func (bc *BlockChain) AddBlock(data string) {
	// 当前区块链是数组形式，上一个区块即为数组最后一个元素
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	// 创建新区块
	newBlock := block.NewBlock(data, prevBlock.Hash)
	// 将新区块加入区块链
	bc.Blocks = append(bc.Blocks, newBlock)
}
