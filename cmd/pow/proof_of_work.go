package main

import (
	"math/big"
	"math"
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
	target.Lsh(target, 256 - targetBits)

	pow := &ProofOfWork{
		block: block,
		target: target,
	}

	return pow
}

func (p *ProofOfWork) Run() (int, []byte) {
	nonce := 1
	var hash []byte

	for nonce < maxNonce {
		p.prepareData(nonce)


		nonce++
	}

	return nonce, hash
}

func (p *ProofOfWork) prepareData(nonce int) {
	//p.block.Nonce = nonce
}