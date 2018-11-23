package main

import (
	"time"
	"bytes"
	"github.com/davecgh/go-spew/spew"
)

type Block struct {
	Timestamp int64
	Data      []byte
	PrevBlockHash []byte
	Nonce     int64
	Hash      []byte
}


func (b *Block) SetHash() {
	b.Hash = bytes.Join([][]byte{

		//b.Timestamp,
	}, []byte{})
}


func NewBlock(data string, prevBlockHash []byte) *Block {
	t, _ := time.Parse(time.RFC3339, "2018-11-23T05:30:00+09:00")
	block := &Block {
		Timestamp: t.Unix(),
		Data: []byte(data),
		PrevBlockHash: prevBlockHash,
	}

	// Pow
	pow:= NewProofOfWork(block)
	//spew.Dump(pow)
	//pow.
	nonce, hash := pow.Run()
	block.Nonce = nonce
	block.Hash = hash

	spew.Dump(block)
	//nonce, hash := NewProofOfWork()
	//block.Hash = hash
	//block.Nonce = nonce
	// Set nance
	// Set Hash


	return block
}


func NewGenesisBlock() *Block {
	return NewBlock("Genesis B lock", nil)
}
