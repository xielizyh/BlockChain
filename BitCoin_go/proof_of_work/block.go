package main

import (
	"time"
)

// Block 由区块头和交易组成
type Block struct {
	Timestamp     int64
	PrevBlockHash []byte
	Hash          []byte
	Data          []byte
	Nonce         int
}

// NewBlock 用于生成新区块，需要运行工作量证明找到有效哈希
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		Data:          []byte(data),
		Nonce:         0,
	}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	// block.Hash = hash
	block.Nonce = nonce

	return block
}

// NewGenesisBlock 生成创世区块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
