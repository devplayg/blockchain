package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"time"
)

var (
	maxNonce = math.MaxInt64
)

const targetBits = 24 // 3 bytes(first 000000)

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, 256-targetBits)

	pow := &ProofOfWork{
		block:  block,
		target: target,
	}

	return pow
}

func (p *ProofOfWork) Run() ([]byte, int, time.Duration) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 26168650 // 26268650

	t := time.Now()
	for nonce < maxNonce {
		data := p.prepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])
		fmt.Printf("\r%x", hash)

		if hashInt.Cmp(p.target) == -1 {
			break
		} else {
			nonce++
		}
		break
	}

	dur := time.Since(t)
	return hash[:], nonce, dur
}

func (p *ProofOfWork) prepareData(nonce int) []byte {
	//spew.Printf("%x / %x / %d / %d / %d\n", p.block.PrevBlockHash, p.block.Data, p.block.Timestamp, targetBits, nonce)

	b := bytes.Join([][]byte{
		p.block.PrevBlockHash,
		p.block.Data,
		Int64ToHex(p.block.Timestamp),
		Int64ToHex(int64(targetBits)),
		Int64ToHex(int64(nonce)),
	}, []byte{})

	return b
}

func (p *ProofOfWork) Validate() bool {
	var hashInt big.Int
	data := p.prepareData(p.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	//spew.Dump(hashInt)
	return hashInt.Cmp(p.target) == -1

	//var hashInt big.Int
	//
	//data := p.prepareData(p.block.Nonce)
	//hash := sha256.Sum256(data)
	//hashInt.SetBytes(hash[:])
	//
	//isValid := hashInt.Cmp(p.target) == -1
	//
	//return isValid
}
