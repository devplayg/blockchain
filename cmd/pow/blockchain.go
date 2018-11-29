package main

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	bc := Blockchain{Blocks: []*Block{}}
	bc.Blocks = append(bc.Blocks, NewGenesisBlock())
	return &bc
}
