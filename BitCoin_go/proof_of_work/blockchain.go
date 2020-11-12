package main

// BlockChain 是一个Block指针数组
type BlockChain struct {
	blocks []*Block
}

// NewBlockChain 创建一个区块链
func NewBlockChain() *BlockChain {
	// 强转为指针数组
	// return &BlockChain{[]*Block{NewGenesisBlock()}}
	return &BlockChain{blocks: []*Block{NewGenesisBlock()}}
}

// AddBlock 向链中加入一个新块
func (bc *BlockChain) AddBlock(data string) {
	// 当前最后一个区块
	preBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, preBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}
