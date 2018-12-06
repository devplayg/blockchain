package main

import (
	"time"
	"bytes"
	"encoding/gob"
	log "github.com/sirupsen/logrus"
)

type Block struct {
	Timestamp int64
	Data      []byte
	PrevBlockHash []byte
	Nonce     int
	Hash      []byte
	dur time.Duration
}

func (b *Block) Serialize() []byte {
	//var network bytes.Buffer        // Stand-in for a network connection
	//enc := gob.NewEncoder(&network) // Will write to network.
	//dec := gob.NewDecoder(&network) // Will read from network.
	//err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	//if err != nil {
	//	log.Fatal("encode error:", err)
	//}
	// Decode (receive) the value.
	//var q Q
	//err = dec.Decode(&q)
	//var buf []byte
	//w := buffer.NewWriter(buf)
	//gob.NewEncoder(w)

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(b)
	if err != nil {
		log.Error(err)
	}

	return buf.Bytes()
}

func DeserializeBlock(data []byte) *Block {
	var block Block
	dec := gob.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(&block)
	if err != nil {
		log.Error(err)
	}
	return &block
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
	hash, nonce, dur := pow.Run() // 26268650
	block.Hash = hash
	block.dur = dur
	block.Nonce = nonce // 11304936

	return block
}


func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", nil)
}
