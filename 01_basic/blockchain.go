package main


type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	lastBlock := bc.Blocks[len(bc.Blocks)-1]
	block := NewBlock(data, lastBlock.Hash)
	bc.Blocks = append(bc.Blocks, block)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}

}
