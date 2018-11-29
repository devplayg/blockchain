package objs

import (
	"bytes"
	"strconv"
	"time"
	"crypto/sha256"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	Hash          []byte
	PrevBlockHash []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	data := bytes.Join([][]byte{timestamp, b.PrevBlockHash, b.Data}, []byte{})
	hash := sha256.Sum256(data)
	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		PrevBlockHash: prevBlockHash,
		Data:          []byte(data),
		Timestamp:     time.Now().Unix(),
		Hash:          []byte{},
	}
	block.SetHash()
	return &block
}

func NewGenesisBlock() *Block {
	block := NewBlock("Genesis block", nil)
	return block
}