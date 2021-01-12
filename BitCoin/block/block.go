package block

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block 区块
type Block struct {
	Timestamp    int64  // 时间戳
	Data         []byte // 交易数据
	PreBlockHash []byte // 父哈希
	Hash         []byte // 当前块哈希
}

// SetHash 计算并设置当前区块的hash
func (b *Block) SetHash() {
	// 将时间戳转换为字节切片
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	// https://studygolang.com/articles/6572
	// Join函数的功能是用字节切片sep把s中的每个字节切片连成一个字节切片并返回
	headers := bytes.Join([][]byte{b.PreBlockHash, b.Data, timestamp}, []byte{})
	// 求区块头的sha256值
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// NewBlock 创建区块
func NewBlock(data string, preBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), preBlockHash, []byte{}}
	block.SetHash()
	return block
}

// NewGenesisBlock 创建创世区块
func NewGenesisBlock() *Block {
	// 创世区块没有父哈希
	return NewBlock("Genesis Block", []byte{})
}
