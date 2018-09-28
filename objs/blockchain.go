package objs

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	lastBlockHash := bc.Blocks[len(bc.Blocks)-1]
	block := NewBlock(data, lastBlockHash.Hash)
	bc.Blocks = append(bc.Blocks, block)
}

func NewBlockchain() *Blockchain {
	blocks := []*Block{NewGenesisBlock()}
	return &Blockchain{blocks}
}
