package main

import (
	"bytes"
	"encoding/gob"
	"log"
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

// Serialize 将Block序列化一个字节数组
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	encoder.Encode(b)

	return result.Bytes()
}

// Deserialize 将字节数组反序列化一个 Block
func Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
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
