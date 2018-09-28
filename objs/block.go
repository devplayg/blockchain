package objs

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
	}
	block.SetHash()
	return &block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	header := bytes.Join([][]byte{b.Hash, b.PrevBlockHash, timestamp}, []byte{})
	hash := sha256.Sum256(header)
	b.Hash = hash[:]
}
