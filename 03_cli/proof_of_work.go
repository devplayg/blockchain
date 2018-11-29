package main

import (
	"math/big"
	"math"
	"bytes"
	"crypto/sha256"
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

func (p *ProofOfWork) Run() ([]byte, int) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	for nonce < maxNonce {
		data := p.prepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		//fmt.Printf("\r%x", hash)

		if hashInt.Cmp(p.target) == -1 {
			break
		} else {
			nonce++
		}

		//break

		//p.Run()
		//hash, nonce := p.Run()
		//if p.isValidBlock() {
		//	break
		//} else {
		//	nonce++
		//}
		//nonce++
		//time.Sleep(1500 * time.Millisecond)
	}

	return hash[:], nonce
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

func (p *ProofOfWork) isValidBlock() bool {
	return true
}